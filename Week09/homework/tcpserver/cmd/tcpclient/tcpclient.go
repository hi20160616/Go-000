package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"golang.org/x/sync/errgroup"
)

const (
	address = ":12345"
)

func connWrite(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	sendChan := make(chan string)
	reciveChan := make(chan string)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	// Deadline
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		select {
		case sig := <-sigs:
			fmt.Println()
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			defer cancel()
			return nil
		case <-ctx.Done():
			defer cancel()
			fmt.Println("client sending routine stoped.")
			return ctx.Err()
		}
	})

	// Send
	g.Go(func() error {
		for {
			fmt.Print(">> ")
			input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			sendChan <- input
			switch strings.TrimSpace(string(input)) {
			case "STOP":
				fmt.Println("Send stop signal to TCP Server...")
				sendChan <- "0"
			case "RESTART":
				fmt.Println("Send restart signal to TCP Server...")
				sendChan <- "6"
			case "EXIT":
				fmt.Println("TCP Client exiting...")
				return nil
			}
			connWrite(conn, sendChan)
		}
	})

	// Recive
	g.Go(func() error {
		for {
			recive, _ := bufio.NewReader(conn).ReadString('\n')
			reciveChan <- recive
			for msg := range reciveChan {
				fmt.Print("->: ")
				fmt.Println(msg)
			}
		}
	})

	if err := g.Wait(); err != nil {
		log.Printf("tcpclient main error: %v", err)
	}
}
