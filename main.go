package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

func main() {
	r := mux.NewRouter()

	// mock data
	books = append(books, Book{ID: "1", Isbn: "148073537", Title: "Intro to HLF", Author: &Author{FirstName: "Delo", LastName: "Delo"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
