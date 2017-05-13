package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Message struct {
	Date            time.Time
	Message, Author string
}

var serverUrl, _ = url.Parse("http://127.0.0.1:8080")

func main() {
	newMsg := Message{
		time.Now(),
		"What",
		"Who"}

	// Get previous messages from server
	index()
	var err error = nil
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	auth, _ := reader.ReadString('\n')

	newMsg.Author = strings.Trim(auth, "\n")

	for err == nil {
		fmt.Print("Enter text: ")
		message, _ := reader.ReadString('\n')
		newMsg.Date = time.Now()
		newMsg.Message = strings.Trim(message, "\n")
		create(newMsg)
		index()
	}
	index()
}

func index() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", serverUrl.String()+"/messages", nil)

	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	v := []Message{}
	err = decoder.Decode(&v)
	if err != nil {
		log.Fatalln(err)
	}
	displayMessage(v)
}

func create(newMsg Message) {
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(newMsg)

	res, err := http.Post(
		serverUrl.String()+"/messages/new",
		"application/json",
		body)

	if err != nil {
		log.Fatal(err)
	} else {
		defer res.Body.Close()
		_, err := io.Copy(os.Stdout, res.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func displayMessage(msgArr []Message) {
	for _, msg := range msgArr {
		fmt.Printf("%s - %s: %s\n", msg.Date, msg.Author, msg.Message)
	}
}
