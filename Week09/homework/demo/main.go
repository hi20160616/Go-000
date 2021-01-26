package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func send(conn net.Conn, ch <-chan string) {
	for e := range ch {
		fmt.Fprintln(conn, e)
	}
}

func recive(conn net.Conn) {
	sendChan := make(chan string, 8)
	defer conn.Close()
	go send(conn, sendChan)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Print("->: ")
		fmt.Println(scanner.Text())
		sendChan <- "Server recive your msg: " + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Println("Scan conn error: ", err)
	}
}

func main() {
	l, err := net.Listen("tcp", ":123")
	if err != nil {
		log.Println(err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go recive(conn)

	}
}
