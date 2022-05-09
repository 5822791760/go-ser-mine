package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

//==================================

func BookByID(id string) (*book, error) {
	for k, v := range books {
		if id == v.ID {
			return &books[k], nil
		}
	}
	return nil, errors.New("Book Not Found")
}

func HelloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO")
}

func FormFunc(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Unable to ParseForm Error err:%s", err)
		return
	}
	fmt.Fprintln(w, "Form Request Successful")
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")

	fmt.Fprintf(w, "First Name = %s\n", fname)
	fmt.Fprintf(w, "Last Name = %s\n", lname)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
	return
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]

	book, err := BookByID(id)
	if err != nil {
		fmt.Fprintf(w, "Error %s", err)
	}
	json.NewEncoder(w).Encode(*book)
}

//==================================
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", HelloFunc)
	r.HandleFunc("/form", FormFunc)
	r.HandleFunc("/books", GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", GetBookByID).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	fmt.Println("Starting Server at Port :8080")
	http.ListenAndServe(":8080", r)
}
