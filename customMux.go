package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type CustomServeMux struct {
}

func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandomFloat(w, r)
		return
	}

	if r.URL.Path == "/randomFloat" {
		giveRandomFloat(w, r)
		return
	}

	if r.URL.Path == "/randomInt" {
		giveRandomInt(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func giveRandomFloat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "your random float number is: %f", rand.Float64())
}

func giveRandomInt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "your random integer number is: %d", rand.Intn(100))
}

func main() {
	mux := &CustomServeMux{}
	http.ListenAndServe(":8008", mux)
}
