package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"cjdavis.me/elysium/models"
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
	p := getProfile()

	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getCurrentOccupation() *models.Occupation {
	return &models.Occupation{
		CompanyName: "CQL, Inc.",
		Position:    "Software Developer",
	}
}

func getProfile() *models.Profile {
	return &models.Profile{
		FirstName:  "CJ",
		LastName:   "Davis",
		Age:        26,
		Occupation: getCurrentOccupation(),
	}
}
