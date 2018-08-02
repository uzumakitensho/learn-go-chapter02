package main

import (
	"io"
	"log"
	"net/http"
)

func MyServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/", MyServer)
	http.HandleFunc("/hello", MyServer)
	log.Fatal(http.ListenAndServe(":8008", nil))
}
