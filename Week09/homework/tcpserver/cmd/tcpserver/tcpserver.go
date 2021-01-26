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
	address string
}

func (s *Server) Start(ctx context.Context) error {
	l, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		return err
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			return err
		}
		fmt.Print("-> ", string(netData))
		t := time.Now()
		msg := "Server time: " + t.Format(time.RFC3339) + "\n" // msg prepare
		c.Write([]byte(msg))                                   // send msg to client
	}
}

func (s *Server) Stop(ctx context.Context) error {
	log.Println("tcp server stopped.")
	return nil
}

func main() {
	s := &Server{address: address}
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

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
			fmt.Println()
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			defer cancel()
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
