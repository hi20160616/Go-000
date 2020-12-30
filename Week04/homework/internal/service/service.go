package service

import (
	"context"
	"strconv"

	pb "github.com/hi20160616/Go-000/Week04/homework/api/helloworld/v1"
	"github.com/hi20160616/Go-000/Week04/homework/internal/biz"
)

const port = ":50051"

type Server struct {
	gc *biz.GreeterCase
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// dto -> do
	h := &biz.Greeter{Name: in.Name, Msg: in.Msg}

	// call biz
	s.gc.SetID(h)

	return &pb.HelloReply{Message: "hello " + in.GetName() + ", Server recived your message: " + in.GetMsg() + " ID: " + strconv.Itoa(int(h.ID))}, nil
}

func NewGreeterServer(gc *biz.GreeterCase) pb.GreeterServer {
	return &Server{gc: gc}
}
