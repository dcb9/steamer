package youku

import (
	"testing"

	"github.com/dcb9/steamer/task/common"
)

func TestFetchInfo(t *testing.T) {
	wantStreams := make(map[string]common.Stream)
	youkuStreams := map[string]YoukuStream{
		"flvhd": {"flvhd", "flv", "标清", 28090073},
		"hd2":   {"hd2", "flv", "超清", 117577724},
		"mp4":   {"mp4", "mp4", "高清", 53866594},
	}
	for key, stream := range youkuStreams {
		wantStreams[key] = stream
	}

	cases := []struct {
		in   string
		want common.ResourceInfo
	}{
		{
			"http://v.youku.com/v_show/id_XMTM4MzA1NDUxNg==.html",
			common.ResourceInfo{
				0,
				"优酷 (Youku)",
				"魔术《心想成真》 刘谦 Mirko Callaci 24",
				"http://v.youku.com/v_show/id_XMTM4MzA1NDUxNg==.html",
				wantStreams,
			},
		},
	}

	for _, c := range cases {
		got := fetchInfo(c.in)
		want := c.want
		if got.Title != want.Title {
			t.Errorf("title did not match got: %s want: %s", got.Title, want.Title)
		}

		if got.Site != want.Site {
			t.Errorf("site did not match got: %s want: %s", got.Site, want.Site)
		}

		if got.Url != want.Url {
			t.Errorf("url did not match got: %s want: %s", got.Url, want.Url)
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
}
