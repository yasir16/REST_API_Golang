package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var book []Book

func main() {
	router := mux.NewRouter()

	book = append(book, Book{ID: 1, Title: "Belonging a things", Author: "Mr. P", Year: "2019"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets all books")
	json.NewEncoder(w).Encode(book)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets one books")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add books")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update books")
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove books")
}
