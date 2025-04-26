package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/davidreque/rest-api-with-go/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	// read dynamic parameter from URL
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"]) 
	
	// iterate over all the mock books and find the book with the same id
	for _, book := range mocks.Books {
		if book.Id == id {
			// if ids are equal, return the book as a JSON response
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			break
		}
	}

}