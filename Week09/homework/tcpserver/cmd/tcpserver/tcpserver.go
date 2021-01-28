package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	address = ":12345"
)

var trans = make(chan string)

type Server struct {
	address  string
	listener net.Listener
}

func (s *Server) Start(ctx context.Context) error {
	l, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	s.listener = l
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		defer conn.Close()
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			text := scanner.Text()
			fmt.Println("-> ", text)
			t := time.Now()
			msg := "Server time: " + t.Format(time.RFC3339) + "\n" // msg prepare
			switch text {
			case "init 0":
				msg = msg + "recive command: STOP"
				conn.Write([]byte(msg))
				fmt.Println(msg)
				s.Stop(ctx)
			case "init 6":
				msg = msg + "recive command: RESTART"
				conn.Write([]byte(msg))
				fmt.Println(msg)
				s.Restart(ctx)
			default:
				conn.Write([]byte(msg)) // send msg to client
			}
		}
	}
}

func (s *Server) Stop(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	log.Println("tcp server stop now...")
	defer cancel()
	s.listener.Close()
	os.Exit(0)
	return ctx.Err()
}

func (s *Server) Restart(ctx context.Context) error {
	fmt.Println("Stop tcp server...")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	s.listener.Close()
	fmt.Println("Start tcp server...")
	if err := s.Start(ctx); err != nil {
		return err
	}
	fmt.Println("Restart tcp server success")
	return ctx.Err()
}

func main() {
	s := &Server{address: address}
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	// Serve
	g.Go(func() error {
		defer cancel()
		fmt.Println("Hi there, server working at ", s.address)
		return s.Start(ctx)
	})

	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		select {
		case sig := <-sigs:
			defer cancel()
			fmt.Println()
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			s.Stop(ctx)
			return nil
		case <-ctx.Done():
			defer cancel()
			s.Stop(ctx)
			return ctx.Err()
		}
	})
	if err := g.Wait(); err != nil {
		log.Printf("tcpserver main error: %v", err)
	}
}
