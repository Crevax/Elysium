package db

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

var appdb *mgo.Session
var Database = os.Getenv("MONGO_DB_NAME")

func init() {
	s, err := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	if err != nil {
		log.Fatalf("Error connecting to the database: " + err.Error())
	}
	s.SetMode(mgo.Monotonic, true)
	appdb = s
}

func AppDB() *mgo.Session {
	return appdb
}
