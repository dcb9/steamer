package httpHandler

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Fatal("Parse template error", err)
	}

	t.Execute(w, req.Host)
}
