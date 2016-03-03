package controllers

import "github.com/julienschmidt/httprouter"

func Init(router *httprouter.Router) {
	new(ProfileController).Init(router)
}
