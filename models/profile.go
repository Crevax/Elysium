package models

import "gopkg.in/mgo.v2/bson"

type Profile struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	FirstName  string        `bson:"FirstName"`
	LastName   string        `bson:"LastName"`
	Age        int           `bson:"Age"`
	Occupation *Occupation   `bson:"Occupation"`
}
