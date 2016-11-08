package youtube

type YoutubeStream struct {
	Itag      string
	Container string
	Quality   string
	Size      uint64
}

func (s YoutubeStream) DlWith() string {
	return "--itag=" + s.Itag
}
