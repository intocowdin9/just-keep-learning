package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 style='color:steelblue'>Hello</h1>"))
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	book := Book{
		Title:  "The Gunslinger",
		Author: "Stephen King",
		Pages:  304,
	}

	json.NewEncoder(w).Encode(book)
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/book", GetBook)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
