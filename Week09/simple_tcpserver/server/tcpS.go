package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + args[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	rd := bufio.NewReader(c)
	wr := bufio.NewWriter(c)

	for {
		recive, _, err := rd.ReadLine()
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(recive)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Println("-> ", string(recive))
		wr.WriteString(time.Now().Format(time.RFC3339) + "\n")
		wr.Flush()
	}
}
