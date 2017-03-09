package main

import (
	"golang.org/x/net/websocket"
	"net/http"

	p "github.com/manuviswam/logmonitor/processor"
	r "github.com/manuviswam/logmonitor/reader"
)

func main() {

	inputChan := make(chan string)
	go r.ReadInfinetly("/Users/manuvisw/file.log", inputChan)
	socketHandler := p.Process(inputChan)

	http.Handle("/logs", websocket.Handler(socketHandler))
	http.Handle("/", http.FileServer(http.Dir("/Users/manuvisw/go/src/github.com/manuviswam/logmonitor/static")))

	panic(http.ListenAndServe(":8080", nil))
}
