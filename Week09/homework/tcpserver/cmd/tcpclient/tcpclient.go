package main

import (
	"bufio"
	"context"
	"errors"
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
	sendChan := make(chan string, 1)
	reciveChan := make(chan string, 1)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	fmt.Println("Welcome to my dark side...")
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		select {
		case sig := <-sigs:
			defer cancel()
			fmt.Println()
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			os.Exit(0)
			return nil
		case <-ctx.Done():
			defer cancel()
			fmt.Println("client routine stoped.")
			return ctx.Err()
		}
	})

	// Send
	g.Go(func() error {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			switch strings.TrimSpace(string(input)) {
			case "STOP":
				fmt.Println("Send stop signal to TCP Server...")
				sendChan <- "init 0\n"
			case "RESTART":
				fmt.Println("Send restart signal to TCP Server...")
				sendChan <- "init 6\n"
			case "EXIT":
				fmt.Println("TCP Client exiting...")
				os.Exit(0)
				return nil
			default:
				sendChan <- input
			}
			fmt.Fprintf(conn, <-sendChan)
		}
	})

	// Recive
	g.Go(func() error {
		for {
			recive, _ := bufio.NewReader(conn).ReadString('\n')
			if recive != "" {
				reciveChan <- recive
				fmt.Print("->: " + <-reciveChan)
			} else {
				return errors.New("recive nothing! maybe server is disconnected...")
			}

		}
	})

	if err := g.Wait(); err != nil {
		log.Printf("tcpclient main error: %v", err)
	}
}
