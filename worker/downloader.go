package worker

import (
	"log"

	"github.com/dcb9/steamer/task/youku"
	"github.com/dcb9/steamer/task/youtube"
	"github.com/dcb9/steamer/util/youget"
	"github.com/gorilla/websocket"
	"encoding/json"
	"strings"
)

func DoDownload(d *DownloadHub) {
	for {
		task := <-d.Add

		info := task.Entity.Info
		if stream, ok := info.Streams[task.Entity.SIndex]; ok {

			options := []string{
				stream.DlWith(),
				"--no-caption",
				"--output-dir",
				d.OutputDir,
				"--output-filename",
				info.Title + ".mp4",
			}

			if youku.IsYouku(info.Url) {
				// do something for youku
			} else if youtube.IsYoutube(info.Url) {
				options = append(options, "-s", "127.0.0.1:1080")
			} else {
				log.Fatalf(`Unknow platform in task: %v`, task)
			}
			options = append(options, info.Url)
			go func() {
				cmd := youget.Cmd(options...)
				stdout, err := cmd.StdoutPipe()
				if err != nil {
					log.Fatal(err)
				}
				stderr, err := cmd.StderrPipe()
				_ = stderr
				if err != nil {
					log.Fatal(err)
				}
				cmd.Start()

				for {
					p := make([]byte, 512)
					n, _ := stdout.Read(p)
					if n <= 0 {
						break;
					}

					for _, listener := range task.Listeners {
						res := struct {
							Callback string `json:"callback"`
							TaskId   int64 `json:"task_id"`
							Data     string `json:"data"`
						}{
							Callback: "downloadProgress",
							TaskId: task.Entity.Id,
							Data: strings.Replace(string(p[:n]), "\n", "\r\n", -1),
						}
						resB, _ := json.Marshal(res)
						if err := listener.WebSocketConn.WriteMessage(websocket.TextMessage, resB); err != nil {
							log.Fatal(err)
						}
					}
				}

				if err := cmd.Wait(); err != nil {
					log.Fatal(err)
				}

				d.Finished <- task
			}()
		} else {
			log.Fatalf(`Unknow stream index in task: %q`, task)
		}
	}
}

func FinishedDownload(d *DownloadHub) {
	for {
		task := <-d.Finished
		go func() {
			for _, listener := range task.Listeners {
				info := task.Entity.Info;
				conn := (*listener).WebSocketConn
				if conn == nil {
					continue
				}

				videoUrl := "/r/" + info.Title + ".mp4"

				res := struct {
					Callback string `json:"callback"`
					TaskId   int64 `json:"task_id"`
					VideoUrl string `json:"video_url"`
				}{
					Callback: "downloaded",
					TaskId: task.Entity.Id,
					VideoUrl: videoUrl,
				}

				resB, _ := json.Marshal(res)
				if err := conn.WriteMessage(websocket.TextMessage, resB); err != nil {
					log.Fatal(err)
				}
			}
		}()
	}
}
