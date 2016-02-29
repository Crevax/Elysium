package db

import (
  "gopkg.in/mgo.v2"
  "os"
  "log"
)

var appdb *mgo.Session
var Database = os.Getenv("MONGOLAB_URI")

func init() {
	s, err := mgo.Dial(os.Getenv("MONGO_DB_NAME"))
	if err != nil {
		log.Println("Error connecting to the database: " + err.Error())
	}
	s.SetMode(mgo.Monotonic, true)
	appdb = s
}

func AppDB() *mgo.Session {
	return appdb
}
