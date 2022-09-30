package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/abba5/grpc-unary-vs-http1-1/model"
)

var c http.Client

func TestMain(m *testing.M) {
	c = http.Client{}
	os.Exit(m.Run())
}

func Test_Someting(t *testing.T) {
	request := model.Request{
		Tid: "tid",
		Sid: "sid",
	}

	b, _ := json.Marshal(&request)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8081/getsession", bytes.NewBuffer(b))

	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("http error %v", err)
	}

	defer resp.Body.Close()
	b, _ = ioutil.ReadAll(resp.Body)

	var response model.Response
	json.Unmarshal(b, &response)

	log.Printf("response %v", response)
}

func Benchmark_grpc_some(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		request := model.Request{
			Tid: "tid",
			Sid: "sid",
		}

		b, _ := json.Marshal(&request)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:8081/getsession", bytes.NewBuffer(b))

		resp, err := c.Do(req)
		if err != nil {
			log.Fatalf("http error %v", err)
		}

		defer resp.Body.Close()
		b, _ = ioutil.ReadAll(resp.Body)

		var response model.Response
		json.Unmarshal(b, &response)
	}
}
