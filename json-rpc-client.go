package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:18888")
	if err != nil {
		panic(err)
	}
	var result int
	args := &Args{4, 5}
	err = client.Call("Calculator.Multiply", args, &result)
	if err != nil {
		panic(err)
	}
	log.Printf("4 * 5 = %v\n", result)
}
