package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/abba5/grpc-unary-vs-http1-1/model"
)

func getSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		jsonRequest, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "body error", http.StatusBadRequest)
			return
		}

		var request model.Request
		err = json.Unmarshal(jsonRequest, &request)
		if err != nil {
			http.Error(w, "request error", http.StatusBadRequest)
			return
		}

		resp := model.Response{
			Tid: "one",
			Sid: "two",
		}

		jsonResp, _ := json.Marshal(&resp)

		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonResp)
		return
	}
}

func main() {
	http.HandleFunc("/getsession", getSession())

	log.Printf("http server start at :%d", 8081)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("can't listen http service %v", err)
	}
}
