package model

type Request struct {
	Tid string `json:"tid"`
	Sid string `json:"sid"`
}

type Response struct {
	Tid string `json:"tid"`
	Sid string `json:"sid"`
}
