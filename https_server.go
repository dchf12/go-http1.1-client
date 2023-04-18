package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func handlerUpgrade(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Connection") == "Upgrade" || r.Header.Get("Upgrade") == "MyProtocol" {
		w.WriteHeader(400)
		return
	}
	fmt.Println("Upgrade to MyProtocol")

	// 低層のソケットを取得
	hijacker := w.(http.Hijacker)
	conn, readWriter, err := hijacker.Hijack()
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	//レスポンスを送信
	response := http.Response{
		StatusCode: 101,
		Header:     make(http.Header),
	}
	response.Header.Set("Upgrade", "MyProtocol")
	response.Header.Set("Connection", "Upgrade")
	response.Write(conn)

	// 通信開始
	for i := 0; i < 10; i++ {
		fmt.Fprintf(readWriter, "%v\n", i)
		fmt.Println("->", i)
		readWriter.Flush()
		recv, err := readWriter.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("<- %v", string(recv))
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	http.HandleFunc("/", handlerUpgrade)
	log.Println("start http listening :18443")
	err := http.ListenAndServeTLS("localhost:18443", "server.crt", "server.key", nil)
	log.Println(err)
}
