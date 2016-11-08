package youku

import (
	"encoding/json"
	"log"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/dcb9/steamer/task/common"
	"github.com/dcb9/steamer/util/youget"
)

func IsYouku(url string) bool {
	return strings.Contains(url, "youku.com")
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
	stdOut, _ := youget.Excecute("--debug", "--json", "-i", url)

	js, err := simplejson.NewJson(stdOut.Bytes())
	if err != nil {
		log.Fatal("New json error: ", err)
	}

	return composeInfo(js)
}

func composeInfo(js *simplejson.Json) common.ResourceInfo {
	var r common.ResourceInfo

	r.Title, _ = js.Get("title").String()
	r.Site, _ = js.Get("site").String()
	r.Url, _ = js.Get("url").String()

	r.Streams = streamsUnmarshal(js)

	return r
}

func streamsUnmarshal(js *simplejson.Json) map[string]common.Stream {
	rawStreams, _ := js.Get("streams").Map()

	streams := make(map[string]common.Stream)
	var stream YoukuStream

	for key := range rawStreams {
		val, _ := js.Get("streams").Get(key).Map()
		stream.Format = key
		for k, v := range val {
			switch k {
			case "container":
				stream.Container = v.(string)
			case "size":
				v := v.(json.Number)
				intV, _ := v.Int64()
				stream.Size = uint64(intV)
			case "video_profile":
				stream.VideoProfile = v.(string)
			}
		}
		streams[key] = stream
	}

	return streams
}
