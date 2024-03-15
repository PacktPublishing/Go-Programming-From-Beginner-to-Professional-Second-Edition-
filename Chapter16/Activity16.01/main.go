package main

import (
	"fmt"
	"log"
	"net/http"
)

type PageWithCounter struct {
	Counter int    `json:"counter"`
	Heading string `json:"heading"`
	Content string `json:"content"`
}

func (h *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Counter++
	msg := fmt.Sprintf("<h1>%s</h1>\n<p>%s<p>\n<p>Views: %d</p>", h.Heading, h.Content, h.Counter)
	w.Write([]byte(msg))
}

func main() {
	hello := PageWithCounter{Heading: "Hello World", Content: "This is the main page"}
	cha1 := PageWithCounter{Heading: "Chapter 1", Content: "This is the first chapter"}
	cha2 := PageWithCounter{Heading: "Chapter 2", Content: "This is the second chapter"}

	http.Handle("/", &hello)
	http.Handle("/chapter1", &cha1)
	http.Handle("/chapter2", &cha2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
