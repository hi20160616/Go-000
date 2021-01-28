package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONN := args[1]
	c, err := net.Dial("tcp", CONN)
	if err != nil {
		fmt.Println(err)
		return
	}

	stdin := bufio.NewReader(os.Stdin)
	rd := bufio.NewReader(c)
	wr := bufio.NewWriter(c)

	for {
		fmt.Print(">> ")
		input, _, _ := stdin.ReadLine()
		wr.Write(input)
		wr.WriteByte('\n')
		wr.Flush()

		recive, _, _ := rd.ReadLine()
		if strings.TrimSpace(string(input)) == "STOP" {
			fmt.Print("TCP client exiting...\n")
			return
		}
		fmt.Printf("->: %s\n", recive)
	}
}
