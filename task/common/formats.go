package common

type ResourceInfo struct {
	Id      int64
	Site    string
	Title   string
	Url     string
	Streams map[string]Stream
}

type Stream interface {
	DlWith() string
}

type DownloadTask struct {
	Id     int64
	Info   *ResourceInfo
	InfoId int64
	SIndex string
	Stdout string
	Stderr string
	Status uint8
}

const (
	sReady   uint8 = 10
	sRunning uint8 = 20
	sTimeout uint8 = 30
	sFinish  uint8 = 40
)
