package main

import (
	"log"
	"net/http"

	"github.com/davidreque/rest-api-with-go/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", handlers.UpadateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods(http.MethodDelete)


	log.Println("Server is running")
	http.ListenAndServe(":4000", router)
}
