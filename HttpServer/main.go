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

const port int = 9001

func main() {

	msg := fmt.Sprintf("Server has started on port: %d\r\n", port)
	log.Print(msg)

	http.HandleFunc("/", homePage)

	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("Fail to start server")
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.RemoteAddr)
}
