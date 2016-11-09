package httpHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/dcb9/steamer/worker"
	"github.com/dcb9/steamer/app"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 10 * time.Second

	// Send pings to peer with this period. WebSocketMust be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 2048
)

func doTicker(ticker *time.Ticker, conn *websocket.Conn) {
	for {
		<-ticker.C
		conn.SetWriteDeadline(time.Now().Add(writeWait))
		if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
			checkErr(err)
		}
	}
}

func initConn(conn *websocket.Conn) {
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
}

func WebSocket(conn *websocket.Conn, w http.ResponseWriter, r *http.Request) {
	ticker := time.NewTicker(pingPeriod)
	initConn(conn)

	outputDir := app.OutputDir
	downloadHub := worker.DownloadHub{
		Add: make(chan *worker.DownloadTask),
		Finished: make(chan *worker.DownloadTask),
		OutputDir: outputDir,
	}

	defer func() {
		log.Println("Websocket connection colsed")
		ticker.Stop()
		conn.Close()
	}()

	go worker.DoDownload(&downloadHub)
	go worker.FinishedDownload(&downloadHub)
	go doTicker(ticker, conn)

	log.Println("Websocket connected")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		if messageType == websocket.TextMessage {
			var r request
			err := json.Unmarshal(p, &r)
			checkErr(err)
			if r.isDownloadRequest() {
				go AddDownloadTask(p, conn, &downloadHub)
			} else if r.isSearchRequest() {
				go Search(p, conn)
			} else {
				return
			}
		}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
