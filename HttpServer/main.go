package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type (

	// Logger logs
	Logger interface {
		Debug(c context.Context, format string, args ...interface{})
	}

	loggerFactory func() Logger
)

var (
	// New creates logger
	New loggerFactory
)

const port int = 9002

func main() {

	log.Printf("Server has started on port: %d\r\n", port)

	http.HandleFunc("/", homePage)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("Fail to start server")
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.RemoteAddr)
}
