package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("/home/hafid/go/src/github.com/uzumakitensho/chapter02/static"))
	log.Fatal(http.ListenAndServe(":8008", router))
}
