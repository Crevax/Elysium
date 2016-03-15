package controllers

import (
	"encoding/json"
	"net/http"

	"cjdavis.me/elysium/services"

	"github.com/julienschmidt/httprouter"
)

type BookController struct{}

func (ctrl *BookController) Init(router *httprouter.Router) {
	router.GET("/book", GetAllBooks)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books := services.GetBookService().GetAllBooks()

	js, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
