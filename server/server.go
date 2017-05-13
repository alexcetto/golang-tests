package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)


type Message struct {
	Date time.Time
	Message, author string
}

var globalMessages = make([]Message, 0)

func main() {
	helloMessage := Message{
		time.Now(),
		"Hello from server",
		"Server"	}

	fmt.Println(helloMessage.Date)

	globalMessages = append(globalMessages, helloMessage)

	connectionArgs := os.Args[1:]
	http.HandleFunc("/messages", allMessagesHandler)
	http.HandleFunc("/messages/new", createMessageHandler)

	var err error = nil
	for err == nil {
		fmt.Printf("Starting server on port %s.\n", connectionArgs[0])

		log.Fatal(http.ListenAndServe(":"+connectionArgs[0], nil))
	}


}

func allMessagesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/messages")
	fmt.Println(globalMessages)
	jsonData, _ := json.Marshal(globalMessages)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func createMessageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/messages/new")

	var newMsg Message

	if r.Body == nil {
		http.Error(w, "Send a request body", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&newMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	globalMessages = append(globalMessages, newMsg)

	fmt.Println(newMsg.Date)
}
