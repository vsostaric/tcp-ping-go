package main

import (
	"log"
	"net"
	"strconv"
	"strings"
)

const (
	message = "Ping"
)

func main() {

	for {
		addr := strings.Join([]string{"localhost", strconv.Itoa(8090)}, ":")
		conn, err := net.Dial("tcp", addr)

		defer conn.Close()

		if err != nil {
			log.Fatalln(err)
		}

		conn.Write([]byte(message))
		conn.Write([]byte("\r\n\r\n"))
		log.Printf("Send: %s", message)

		buff := make([]byte, 1024)
		n, _ := conn.Read(buff)
		log.Printf("Receive: %s", buff[:n])
	}
}
