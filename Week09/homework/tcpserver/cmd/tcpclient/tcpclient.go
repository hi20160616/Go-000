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

func main() {
	inchan := make(chan string)
	outchan := make(chan string)
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
			// TODO: stop action
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
			inchan <- input
		}
	})

	// Recive
	g.Go(func() error {
		for {
			recive, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("->: ")
			outchan <- recive
		}
	})

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		input, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, input+"\n")

		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + msg)
		switch strings.TrimSpace(string(input)) {
		case "EXIT":
			fmt.Println("TCP Client exiting...")
			return
		case "STOP":
			fmt.Println("Send stop signal to TCP Server...")
		case "SLEEP":
			fmt.Println("Send sleep signal to TCP server...")
		}
	}
}
