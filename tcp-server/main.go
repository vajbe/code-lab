package main

import (
	"fmt"
	"log"
	"net"
)

func do(listener net.Listener) {
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}

	readSlice := []byte("")
	conn.Read(readSlice)
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 13\r\n" +
		"\r\n" +
		"Hello, World!"
	conn.Write([]byte(response))

}

func main() {
	listner, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Server started on 8081")
	for {
		do(listner)
	}
	listner.Close()
}
