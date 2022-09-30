package main

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	proto "github.com/abba5/grpc-unary-vs-http1-1/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var c proto.SessionClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c = proto.NewSessionClient(conn)

	os.Exit(m.Run())
}

func Test_Someting(t *testing.T) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ValidateSession(ctx, &proto.Request{
		Tid: "this is tid",
		Sid: "this is sid",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetTid())
	log.Printf("Greeting: %s", r.GetSid())
}

func Benchmark_grpc_some(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, err := c.ValidateSession(ctx, &proto.Request{
			Tid: "this is tid",
			Sid: "this is sid",
		})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	}
}
