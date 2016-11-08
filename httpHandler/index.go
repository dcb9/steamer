package httpHandler

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	data, err := Asset("views/index.html")
	if err != nil {
		log.Fatal("Asset not found")
	}
	t, err := template.New("home").Parse(string(data))
	if err != nil {
		log.Fatal("Parse template error", err)
	}

	t.Execute(w, req.Host)
}
