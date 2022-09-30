package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	proto "github.com/abba5/grpc-unary-vs-http1-1/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedSessionServer
}

func (s *server) ValidateSession(_ context.Context, req *proto.Request) (*proto.Response, error) {
	// log.Printf("tid: %v", req.GetTid())
	_ = time.Now()
	return &proto.Response{
		Tid: "one",
		Sid: "two",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterSessionServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
