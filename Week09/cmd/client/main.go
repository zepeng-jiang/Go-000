package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	addr := &net.TCPAddr{Port: 8000}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalf("dial tcp is failed. error: %v \n", err)
	}
	defer conn.Close()
	ch := make(chan bool)
	go send(conn, ch)
	go read(conn, ch)
	<-ch
}

func send(conn net.Conn, exitCh <-chan bool) {
	r := bufio.NewReader(os.Stdin)
	for {
		log.Print("请输入: ")
		b, err := r.ReadBytes('\n')
		if err != nil {
			log.Printf("read message is failed, error: %v \n", err)
			break
		}
		_, err = conn.Write(b)
		if err != nil {
			log.Printf("write message is failed, error: %v \n", err)
			break
		}
	}
}

func read(conn net.Conn, exitCh chan<- bool) {
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Printf("client read message is failed, error: %v \n", err)
			close(exitCh)
			break
		}
		msg = strings.Replace(msg, "\n", "", -1)
		log.Printf("client receive server message success, message: %v \n", msg)
	}
}