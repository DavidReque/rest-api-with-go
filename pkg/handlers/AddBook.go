package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/davidreque/rest-api-with-go/pkg/mocks"
	"github.com/davidreque/rest-api-with-go/pkg/models"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	// read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// append to the book mocks
	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	// send a 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
