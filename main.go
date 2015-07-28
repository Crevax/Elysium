package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"

	"cjdavis.me/elysium/services"
)

var mongoHost = os.Getenv("MONGOLAB_URI")
var mongoSession *mgo.Session
var profileService services.IProfileService

func init() {
	mongoSession, err := mgo.Dial(mongoHost)
	if err != nil {
		log.Fatalf("Database connection error: %s\n", err.Error())
		os.Exit(1)
	}
	mongoSession.SetMode(mgo.Monotonic, true)
	profileService = services.NewProfileService(mongoSession)
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
	p := profileService.GetProfile()

	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
