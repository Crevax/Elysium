package main

import (
	"log"
	"net/http"
	"os"

	"cjdavis.me/elysium/controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	controllers.Init(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Elysium listening on port %s", port)
	http.ListenAndServe(":"+port, router)
}
