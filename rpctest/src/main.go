package main

import (
	"client"
	"server"
	"time"
)

func main() {
	go server.Server()
	go client.Client()
	time.Sleep(time.Second * 5)
}
