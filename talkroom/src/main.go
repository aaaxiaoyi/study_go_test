package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main()  {
	fmt.Println("start work")
	fmt.Println("......")
	time.Sleep(time.Second)
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("invalid address")
	} else {
		fmt.Println("the address is:", addr.String())
	}
	fmt.Println("stop work")
}
