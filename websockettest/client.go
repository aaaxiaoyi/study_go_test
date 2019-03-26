package main

import (
	"log"
)




func main(){
	ws = new WebSocket("ws://127.0.0.1:8080/echo")
	ws.onopen = function(evt) {...}
	ws.onmessage = function(evt)
}