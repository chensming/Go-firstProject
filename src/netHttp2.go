package main

import (
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("Hello, go"))
}

func main() {
	http.HandleFunc("/hello", SayHello)
	_ = http.ListenAndServe(":8001", nil)

}