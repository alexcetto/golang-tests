package main

import (
	"fmt"
	"net/http"
	"os"
)

type Message struct {
	date, message, author string
}

func main() {
	connectionArgs := os.Args[1:]
	fmt.Printf("Starting server on port %s.\n", connectionArgs[0])
	http.HandleFunc("/", clientHandler)
	http.ListenAndServe(":"+connectionArgs[0], nil)
}

func clientHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Working", r.URL.Path[1:])
}
