package processor

import (
	"fmt"
	"golang.org/x/net/websocket"
)

var connections []*websocket.Conn

var lastTenMsgs []string

func Process(inputChan chan string) func(*websocket.Conn) {
	connections = make([]*websocket.Conn, 0)
	lastTenMsgs = make([]string, 1)
	for msg := range inputChan {
		if msg == "EOF" {
			break
		}
		addToCache(msg)
	}
	return func(ws *websocket.Conn) {
		connections = append(connections, ws)
		for _, line := range lastTenMsgs {
			ws.Write([]byte(line))
		}
		for line := range inputChan {
			broadcast(line)
			addToCache(line)
		}
	}
}

func broadcast(msg string) {
	for _, conn := range connections {
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error ", err)
		}
	}
}

func addToCache(msg string) {
	if len(lastTenMsgs) >= 10 {
		lastTenMsgs = append(lastTenMsgs[:0], lastTenMsgs[1:]...)
	}
	lastTenMsgs = append(lastTenMsgs, msg)
}
