package worker

import (
	"github.com/dcb9/steamer/task/common"
	"github.com/gorilla/websocket"
)

type DownloadListener struct {
	WebSocketConn *websocket.Conn
}

type DownloadTasks map[int64]DownloadTask

type DownloadTask struct {
	Entity    *common.DownloadTask
	Listeners []*DownloadListener
}

type DownloadHub struct {
	Add       chan *DownloadTask
	Finished  chan *DownloadTask
	OutputDir string
}

func (t *DownloadTask) AddListener(l *DownloadListener) {
	t.Listeners = append(t.Listeners, l)
}

func (t *DownloadTask) RemoveListener(l *DownloadListener) {
	for index, listener := range t.Listeners {
		if listener == l {
			t.Listeners = append(t.Listeners[:index], t.Listeners[index + 1:]...)
			break
		}
	}
}
