package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {

			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	time := time.Now().Format("15:04:05\n\r")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		_, err := io.WriteString(c, time + "Server: " +input.Text())
		if err != nil {
			return
		}
	}
}
