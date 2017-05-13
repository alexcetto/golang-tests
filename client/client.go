package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"net/url"
	"time"
	"encoding/json"
	"bytes"
)

type Message struct {
	Date time.Time
	Message, author string
}


var serverUrl, _ = url.Parse("http://127.0.0.1:8080")

func main() {
	newMsg := Message{
		time.Now(),
		"What",
		"Who"	}
	index()
	create(newMsg)
	index()
}

func index() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", serverUrl.String() + "/messages", nil)

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
	log.Println(v)
}

func create(newMsg Message) {
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(newMsg)

	res, err := http.Post(
		serverUrl.String() + "/messages/new",
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
