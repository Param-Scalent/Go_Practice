package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil))
}
