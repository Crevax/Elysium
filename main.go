package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"cjdavis.me/elysium/services"
)

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("Elysium listening on port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	profileService := services.NewProfileService()
	p := profileService.GetProfile()

	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
