package youtube

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dcb9/steamer/task/common"
	"github.com/dcb9/steamer/util/youget"
)

func IsYoutube(url string) bool {
	return strings.Contains(url, "youtube.com")
}

func FetchInfo(url string) common.ResourceInfo {
	resourceInfo := r.EntityByUrl(url)
	if resourceInfo.Id == 0 {
		resourceInfo = fetchInfo(url)
		go r.AddOne(resourceInfo)
	}

	return resourceInfo
}

func fetchInfo(url string) common.ResourceInfo {
	stdout, _ := youget.Excecute("--debug", "-s", "127.0.0.1:1080", "-i", url)
	info := parse(stdout.String())
	info.Url = url
	return info
}

func parse(stdout string) common.ResourceInfo {
	var data common.ResourceInfo
	lines := strings.Split(stdout, "\n")

	contentLines := make([]string, 0)
	for _, row := range lines {
		row := strings.TrimSpace(row)
		if len(row) < 1 || row[0] == '#' || row[0] == '[' {
			continue
		}
		contentLines = append(contentLines, row)
	}

	_, data.Site = parseRowValue(contentLines[0])
	_, data.Title = parseRowValue(contentLines[1])

	fmt.Printf("%q\n", contentLines[3:])
	l := []map[string]string{}
	for _, row := range contentLines[3:] {
		if row[0] == '-' {
			row = row[2:]
		}
		k, v := parseRowValue(row)
		l = append(l, map[string]string{k: v})
	}

	streams := make(map[string]YoutubeStream)
	i := 0
	len := len(l)
	for {
		if i >= len {
			break
		}
		s := YoutubeStream{}
		for _, index := range []string{"itag", "container", "quality", "size"} {
			if i >= len {
				break
			}
			if v, found := l[i][index]; found {
				switch index {
				case "itag":
					s.Itag = v
				case "container":
					s.Container = v
				case "quality":
					s.Quality = v
				case "size":
					s.Size = 0
					reg := regexp.MustCompile(`\([0-9]+`)
					matchedS := string(reg.Find([]byte(v)))[1:]
					s.Size, _ = strconv.ParseUint(matchedS, 10, 64)
				default:
					continue
				}
				i++
			}
		}
		streams[s.Itag] = s
	}

	data.Streams = make(map[string]common.Stream)
	var stream common.Stream
	for index, s := range streams {
		stream = s
		data.Streams[index] = stream
	}

	return data
}

func parseRowValue(row string) (string, string) {
	tmp := strings.Split(row, ":")
	return tmp[0], strings.TrimSpace(tmp[1])
}
