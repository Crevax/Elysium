package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Profile struct {
	FirstName  string
	LastName   string
	Age        int
	Occupation *Occupation
}

type Occupation struct {
	CompanyName string
	Position    string
}

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

func getCurrentOccupation() *Occupation {
	return &Occupation{"CQL, Inc.", "Software Developer"}
}

func getProfile() *Profile {
	return &Profile{"CJ", "Davis", 26, getCurrentOccupation()}
}
