package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/davidreque/rest-api-with-go/pkg/mocks"
	"github.com/davidreque/rest-api-with-go/pkg/models"
	"github.com/gorilla/mux"
)

func UpadateBook(w http.ResponseWriter, r *http.Request) {
	// read dynamic parameter from URL
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// read the request body
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	// iterate over all the mock books and find the book with the same id
	for index, book := range mocks.Books {
		if book.Id == id {
			// update the book with the new data
			book.Title = updatedBook.Title
			book.Author = updatedBook.Author
			book.Description = updatedBook.Description

			mocks.Books[index] = book

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}

}
