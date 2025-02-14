package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

var templateResponse = "HTTP/1.1 200 OK\r\n" +
	"Content-Type: text/plain\r\n" +
	"Content-Length: 13\r\n" +
	"\r\n" +
	"Hello, World!"

func do(conn net.Conn) {
	buffer := make([]byte, 1024)
	fmt.Println("Processing...")
	time.Sleep(8 * time.Second)
	conn.Read(buffer)
	conn.Write([]byte(templateResponse))
	conn.Close()
	fmt.Println("Request processed...")
}

func main() {
	listner, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started on 8081")

	for {
		fmt.Println("waiting for client to connect")
		conn, err := listner.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Client Connected")
		go do(conn)
	}
}
