package pitcher

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	message = "Ping"
)

var messagesSent = 0
var messagesNotReplied = 0
var speed = 0.0
var maximalTime time.Duration = 0
var fullReplyTime time.Duration = 0

func Pitch() {

	log.Printf("Pitching to port: 8090")

	for {
		go createConnection()
		time.Sleep(5 * time.Second)
	}

}

func printState() {
	log.Printf("Messages send: %d", messagesSent)
	log.Printf("Messages not replied: %d", messagesNotReplied)
	log.Printf("Full reply time: %s", fullReplyTime.String())
	log.Printf("Speed: %f messages per second", speed)
	log.Printf("Maximal time: %s", maximalTime.String())
}

func createConnection() {
	addr := strings.Join([]string{"localhost", strconv.Itoa(8090)}, ":")
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)

	defer conn.Close()

	start := time.Now()
	var elapsed time.Duration = 0

	if err != nil {
		log.Fatalln(err)
	}

	conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	conn.Write([]byte(message))
	conn.Write([]byte("\r\n\r\n"))

	buff := make([]byte, 1024)

	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	_, readErr := conn.Read(buff)
	if os.IsTimeout(readErr) {
		messagesNotReplied++
	} else {
		elapsed = time.Since(start)
		fullReplyTime += elapsed

		go checkMaxTime(elapsed)

	}

	messagesSent++
	speed = float64(messagesSent-messagesNotReplied) / fullReplyTime.Seconds()

	go printState()
}

func checkMaxTime(elapsed time.Duration) {
	if elapsed > maximalTime {
		maximalTime = elapsed
	}
}
