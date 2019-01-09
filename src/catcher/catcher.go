package catcher

import (
	"bufio"
	"log"
	"net"
	"os"
	"strconv"
)

const (
	Message = "Pong"
)

func Catch() {
	port := 8090
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	defer listen.Close()
	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, err)
		os.Exit(1)
	}
	log.Printf("Begin listen port: %d", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handler(conn)
	}

}

func handler(conn net.Conn) {
	defer conn.Close()

	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)

	n, _ := r.Read(buf)
	data := string(buf[:n])
	log.Println("Receive:", data)

	w.Write([]byte(Message))
	w.Flush()
}
