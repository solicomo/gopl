package main

import (
	"io"
	"log"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/hello", handleHello)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
