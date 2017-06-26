package main

import (
	"io"
	"net/http"
	"github.com/bmizerany/pat"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hej, "+req.URL.Query().Get(":name")+"!\n")
}

func main() {
	m := pat.New()

	m.Get("/hello/:name", http.HandlerFunc(HelloServer))

	http.Handle("/", m)
	http.ListenAndServe(":12345", nil)
}