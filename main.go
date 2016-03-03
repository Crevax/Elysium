package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	"cjdavis.me/elysium/services"
)

func main() {
	router := httprouter.New()
	router.GET("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Elysium listening on port %s", port)
	http.ListenAndServe(":"+port, router)
}

func handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p := services.GetProfileService().GetProfile()

	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
