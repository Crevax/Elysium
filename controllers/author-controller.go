package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cjdavis.me/elysium/services"

	"github.com/julienschmidt/httprouter"
)

type AuthorController struct{}

func (ctrl *AuthorController) Init(router *httprouter.Router) {
	router.GET("/author", GetAllAuthors)
	router.GET("/author/:id/book", GetAllBooksForAuthor)
}

func GetAllAuthors(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	authors := services.GetAuthorService().GetAllAuthors()

	js, err := json.Marshal(authors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetAllBooksForAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	authorID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	books := services.GetBookService().GetBooksByAuthor(authorID)

	js, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
