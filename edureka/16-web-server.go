package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", handler2)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to My HomePage \n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World \n")
}
