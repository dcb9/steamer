package httpHandler

import (
	"encoding/json"
	"log"

	"github.com/dcb9/steamer/task"
	"github.com/dcb9/steamer/task/common"
	"github.com/gorilla/websocket"
	"github.com/dcb9/steamer/worker"
)

func AddDownloadTask(p []byte, conn *websocket.Conn, hub *worker.DownloadHub) {
	var dData downloadRequest
	var res response
	err := json.Unmarshal(p, &dData)
	checkErr(err)
	entity := addDownloadTask(dData.Url, dData.Id, dData.SIndex)
	res.Uid = dData.Uid;
	if entity.Id == 0 {
		res.Success = 0;
	} else {
		res.Success = 1;
		res.Data = struct{ TaskId int64 `json:"task_id"` }{entity.Id}

		downloadTask := worker.DownloadTask{Entity: &entity}
		listener := worker.DownloadListener{WebSocketConn: conn}
		downloadTask.AddListener(&listener)

		hub.Add <- &downloadTask
	}
	resStr, _ := json.Marshal(res)
	if err = conn.WriteMessage(websocket.TextMessage, resStr); err != nil {
		log.Fatal(err)
	}
}

func addDownloadTask(url string, id int64, sIndex string) common.DownloadTask {
	return task.AddDownloadTask(url, id, sIndex)
}
