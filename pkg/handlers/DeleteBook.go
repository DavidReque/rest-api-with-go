package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/davidreque/rest-api-with-go/pkg/mocks"
	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// read the book ID from the URL
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// iterate over the books and find the book with the given ID
	for index, book := range mocks.Books {
		if book.Id == id {
			// delete the book and send a response if the book is found
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mocks.Books)
			break
		}
	}

}
