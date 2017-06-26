package main

import (
	"net/http"
	"fmt"
)

type defaultHandler struct {
	greeting string
}

func (handler defaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte (fmt.Sprintf("%v, world!", handler.greeting)))
}

func main() {
	http.Handle("/", &defaultHandler{greeting:"Hello"})
	http.ListenAndServe(":8001", nil)
}
