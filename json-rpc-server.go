package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// メソッドが属す型
type Calculator int

// RPCで外部から呼ばれるメソッド
func (c *Calculator) Multiply(args Args, result *int) error {
	log.Println("Multiply caled: %v, %v\n", args.A, args.B)
	*result = args.A + args.B
	return nil
}

// 外部から呼ばれる時の引数
type Args struct {
	A, B int
}

func main() {
	calculator := new(Calculator)
	server := rpc.NewServer()
	server.Register(calculator)
	http.Handle(rpc.DefaultRPCPath, server)
	log.Println("start htttp listening :18888")
	listener, err := net.Listen("tcp", ":18888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
