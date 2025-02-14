package main

import (
	"fmt"
	"log"
	"net"
)

var templateResponse = "HTTP/1.1 200 OK\r\n" +
	"Content-Type: text/plain\r\n" +
	"Content-Length: 13\r\n" +
	"\r\n" +
	"Hello, World!"

func do(conn net.Conn) {
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	conn.Write([]byte(templateResponse))
	conn.Close()
}

func main() {
	listner, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Server started on 8081")

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatal(err)
		}
		do(conn)
	}
}
