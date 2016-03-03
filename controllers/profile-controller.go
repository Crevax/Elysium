package controllers

import (
	"encoding/json"
	"net/http"

	"cjdavis.me/elysium/services"
	"github.com/julienschmidt/httprouter"
)

type ProfileController struct{}

func (ctrl *ProfileController) Init(router *httprouter.Router) {
	router.GET("/", Index)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p := services.GetProfileService().GetProfile()

	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
