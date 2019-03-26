package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)


var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request)  {
	c, _ := upgrader.Upgrade(w, r, nil)
	defer c.Close()
	for {
		mt, message, _ := c.ReadMessage()
		c.WriteMessage(mt, append([]byte("hello"), message[:]...))
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}

func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
}

