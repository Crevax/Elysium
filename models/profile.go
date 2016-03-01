package models

type Profile struct {
	ID         int         `json:"id"`
	FirstName  string      `json:"FirstName"`
	LastName   string      `json:"LastName"`
	Age        int         `json:"Age"`
	Occupation *Occupation `json:"Occupation"`
}
