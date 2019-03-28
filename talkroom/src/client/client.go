package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal error:%s", err.Error())
		os.Exit(1)
	}
}


func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage：%s\n", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("timestamp"))
	checkError(err)
	conn.CloseWrite()
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))

	_, err = conn.Write([]byte("time"))
	checkError(err)
	conn.CloseWrite()
	result, err = ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
