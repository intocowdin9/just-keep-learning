package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	msgChan    chan string
	whoamiChan chan string
)

func getTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if msgChan != nil {
		msg := time.Now().Format("15:04:05")
		msgChan <- msg
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	if whoamiChan != nil {
		msg := "rafli"
		whoamiChan <- msg
	}
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("client connected")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	msgChan = make(chan string)
	whoamiChan = make(chan string)

	defer func() {
		close(msgChan)
		close(whoamiChan)
		msgChan = nil
		whoamiChan = nil
		fmt.Println("client closed connection")
	}()

	flusher, ok := w.(http.Flusher)
	if !ok {
		fmt.Println("could not init http.Flusher")
	}

	for {
		select {
		case message := <-msgChan:
			fmt.Println("case message... sending message")
			fmt.Println(message)
			fmt.Fprintf(w, "data: %s\n\n", message)
			flusher.Flush()
		case message2 := <-whoamiChan:
			fmt.Println("case message2... sending message")
			fmt.Println(message2)
			fmt.Fprintf(w, "data: %s\n\n", message2)
			flusher.Flush()
		case <-r.Context().Done():
			fmt.Println("client closed connection")
			return
		}
	}
}

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/event", sseHandler)
	router.HandleFunc("/time", getTime)
	router.HandleFunc("/user", getUser)

	log.Fatal(http.ListenAndServe(":9000", router))
}
