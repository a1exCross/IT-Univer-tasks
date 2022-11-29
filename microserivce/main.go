package main

import (
	"microservice/service"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":80", service.NewHandler("AlexZabolotskikh")); err != nil {
		panic(err)
	}
}
