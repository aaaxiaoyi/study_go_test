package client

import (
	"fmt"
	"log"
	"net/rpc"
	"server"
)

func Client() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &server.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith:%d*%d=%d\n", args.A, args.B, reply)

	quotient := new(server.Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done
	fmt.Printf("Arith:%d/%d=%d\n", args.A, args.B, quotient.Quo)
	fmt.Printf("Arith:%d%%%d=%d\n", args.A, args.B, quotient.Rem)
	fmt.Printf("%v %v %v\n", server.Args(replyCall.Args).A, server.Args(replyCall.Args).B, replyCall.Reply)
}
