package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":8888"

func handlerRootEN(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {

	http.HandleFunc("/en", handlerRootEN)

	log.Println("[main] listening on port", port)
	log.Fatalln(http.ListenAndServe(port, nil))
}
