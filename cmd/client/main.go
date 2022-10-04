package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/abba5/grpc-unary-vs-http1-1/model"
	proto "github.com/abba5/grpc-unary-vs-http1-1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func callGRPC() http.HandlerFunc {
	conn, err := grpc.Dial("grpcserver:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := proto.NewSessionClient(conn)

	return func(w http.ResponseWriter, r *http.Request) {
		callTheGRPCServer(c)
		w.Write([]byte("all good"))
	}
}

func callTheGRPCServer(c proto.SessionClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.ValidateSession(ctx, &proto.Request{
		Tid: "tid",
		Sid: "sid",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}

func callTheHTTPServer(c *http.Client) {
	request := model.Request{
		Tid: "tid",
		Sid: "sid",
	}

	b, _ := json.Marshal(&request)
	req, _ := http.NewRequest(
		http.MethodPost,
		"http://httpserver:8081/getsession",
		bytes.NewBuffer(b),
	)

	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("http error %v", err)
	}

	defer resp.Body.Close()
	b, _ = ioutil.ReadAll(resp.Body)

	var response model.Response
	json.Unmarshal(b, &response)
}

func callHTTP() http.HandlerFunc {
	c := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   time.Second,
				KeepAlive: time.Minute,
			}).Dial,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		callTheHTTPServer(c)
		w.Write([]byte("all good http"))
	}
}

func main() {
	http.HandleFunc("/http", callHTTP())
	http.HandleFunc("/grpc", callGRPC())

	log.Printf("http server start at :%d", 8081)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("can't listen http service %v", err)
	}
}
