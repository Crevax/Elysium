package models

type Occupation struct {
	CompanyName string `bson:"CompanyName"`
	Position    string `bson:"Position"`
}
