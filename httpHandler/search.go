package httpHandler

import (
	"encoding/json"
	"log"

	"github.com/dcb9/steamer/task"
	"github.com/gorilla/websocket"
)

func Search(p []byte, conn *websocket.Conn) {
	var sData searchRequest
	err := json.Unmarshal(p, &sData)
	checkErr(err)

	info := task.FetchResourceInfo(sData.Url)
	response := response{sData.Uid, 1, info}
	b, _ := json.Marshal(response)
	if err = conn.WriteMessage(websocket.TextMessage, b); err != nil {
		log.Fatal(err)
	}
}
