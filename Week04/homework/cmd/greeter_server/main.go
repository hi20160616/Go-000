package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/hi20160616/Go-000/Week04/homework/api/helloworld/v1"
	srvHandler "github.com/hi20160616/Go-000/Week04/homework/internal/pkg/service_handler"
	"github.com/hi20160616/Go-000/Week04/homework/internal/service"
	"golang.org/x/sync/errgroup"
)

const (
	address = ":50051"
)

func main() {
	gc := InitGreeterCase()
	service := service.NewGreeterServer(gc)

	s := srvHandler.NewServer(address)
	pb.RegisterGreeterServer(s, service)

	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error { return s.Start(ctx) })

	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-sigs:
			fmt.Println()
			log.Printf("signal caught: %s, reday to quit...", sig.String())
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		log.Printf("greeter_server main error: %v", err)
	}
}
