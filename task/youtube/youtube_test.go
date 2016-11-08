package youtube

import (
	"testing"

	"github.com/dcb9/steamer/task/common"
)

func TestParse(t *testing.T) {
	stdout := `
 site:                YouTube
title:               LAKE LOUISE SNOWBOARDING!
streams:             # Available quality and codecs
    [ DASH ] ____________________________________
    - itag:          137
      container:     mp4
      quality:       1920x1080
      size:          551.5 MiB (578326641 bytes)
    # download-with: you-get --itag=137 [URL]

    - itag:          248
      container:     webm
      quality:       1920x1080
      size:          385.1 MiB (403768757 bytes)
    # download-with: you-get --itag=248 [URL]

    - itag:          136
      container:     mp4
      quality:       1280x720
      size:          277.7 MiB (291216180 bytes)
    # download-with: you-get --itag=136 [URL]

    - itag:          247
      container:     webm
      quality:       1280x720
      size:          223.4 MiB (234266553 bytes)
    # download-with: you-get --itag=247 [URL]

    - itag:          135
      container:     mp4
      quality:       854x480
      size:          154.4 MiB (161878667 bytes)
    # download-with: you-get --itag=135 [URL]

    - itag:          244
      container:     webm
      quality:       854x480
      size:          121.0 MiB (126918200 bytes)
    # download-with: you-get --itag=244 [URL]

    - itag:          134
      container:     mp4
      quality:       640x360
      size:          90.3 MiB (94673271 bytes)
    # download-with: you-get --itag=134 [URL]

    - itag:          243
      container:     webm
      quality:       640x360
      size:          77.5 MiB (81286054 bytes)
    # download-with: you-get --itag=243 [URL]

    - itag:          133
      container:     mp4
      quality:       426x240
      size:          53.5 MiB (56060950 bytes)
    # download-with: you-get --itag=133 [URL]

    - itag:          242
      container:     webm
      quality:       426x240
      size:          53.0 MiB (55608468 bytes)
    # download-with: you-get --itag=242 [URL]

    - itag:          278
      container:     webm
      quality:       256x144
      size:          36.1 MiB (37885636 bytes)
    # download-with: you-get --itag=278 [URL]

    - itag:          160
      container:     mp4
      quality:       256x144
      size:          34.0 MiB (35686918 bytes)
    # download-with: you-get --itag=160 [URL]

    [ DEFAULT ] _________________________________
    - itag:          22
      container:     mp4
      quality:       hd720
      size:          277.5 MiB (290933730 bytes)
    # download-with: you-get --itag=22 [URL]

    - itag:          43
      container:     webm
      quality:       medium
    # download-with: you-get --itag=43 [URL]

    - itag:          18
      container:     mp4
      quality:       medium
    # download-with: you-get --itag=18 [URL]

    - itag:          36
      container:     3gp
      quality:       small
    # download-with: you-get --itag=36 [URL]

    - itag:          17
      container:     3gp
      quality:       small
    # download-with: you-get --itag=17 [URL]
    `

	got := parse(stdout)
	want := dataProvider()

	compare(t, got, want)
}

func dataProvider() common.ResourceInfo {
	url := "https://www.youtube.com/watch?v=YB_Sn_CMHug"
	streams := make(map[string]common.Stream)
	youtubeStreams := map[string]YoutubeStream{
		"137": YoutubeStream{"137", "mp4", "1920x1080", 578326641},
		"248": YoutubeStream{"248", "webm", "1920x1080", 403768757},
		"136": YoutubeStream{"136", "mp4", "1280x720", 291216180},
		"247": YoutubeStream{"247", "webm", "1280x720", 234266553},
		"135": YoutubeStream{"135", "mp4", "854x480", 161878667},
		"244": YoutubeStream{"244", "webm", "854x480", 126918200},
		"134": YoutubeStream{"134", "mp4", "640x360", 94673271},
		"243": YoutubeStream{"243", "webm", "640x360", 81286054},
		"133": YoutubeStream{"133", "mp4", "426x240", 56060950},
		"242": YoutubeStream{"242", "webm", "426x240", 55608468},
		"278": YoutubeStream{"278", "webm", "256x144", 37885636},
		"160": YoutubeStream{"160", "mp4", "256x144", 35686918},
		"22":  YoutubeStream{Itag: "22", Container: "mp4", Quality: "hd720", Size: 290933730},
		"43":  YoutubeStream{Itag: "43", Container: "webm", Quality: "medium", Size: 0},
		"18":  YoutubeStream{Itag: "18", Container: "mp4", Quality: "medium", Size: 0},
		"36":  YoutubeStream{Itag: "36", Container: "3gp", Quality: "small", Size: 0},
		"17":  YoutubeStream{Itag: "17", Container: "3gp", Quality: "small", Size: 0},
	}

	for k, v := range youtubeStreams {
		streams[k] = v
	}

	want := common.ResourceInfo{
		Id:      0,
		Site:    "YouTube",
		Title:   "LAKE LOUISE SNOWBOARDING!",
		Url:     url,
		Streams: streams,
	}
	return want
}

func TestFetchInfo(t *testing.T) {
	want := dataProvider()
	cases := []struct {
		in   string
		want common.ResourceInfo
	}{
		{
			want.Url,
			want,
		},
	}

	for _, c := range cases {
		got := fetchInfo(c.in)
		compare(t, got, c.want)
		if got.Url != c.want.Url {
			t.Errorf("url did not match got: %s want: %s", got.Url, want.Url)
		}
	}
}

func compare(t *testing.T, got common.ResourceInfo, want common.ResourceInfo) {
	if got.Title != want.Title {
		t.Errorf("title did not match got: %s want: %s", got.Title, want.Title)
	}

	if got.Site != want.Site {
		t.Errorf("site did not match got: %s want: %s", got.Site, want.Site)
	}

	if len(got.Streams) != len(want.Streams) {
		t.Errorf("streams did not match got: %s want: %s", got.Streams, want.Streams)
	} else {
		for key, wantStream := range want.Streams {
			gotStream, ok := got.Streams[key]
			if !ok {
				t.Errorf("stream key %s did not exists", key)
				continue
			}

			if wantStream != gotStream {
				t.Errorf("streams[%s] did not match got: %q want: %q", key, gotStream, wantStream)
			}
		}
	}
}
