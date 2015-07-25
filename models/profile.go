package models

type Profile struct {
	FirstName  string
	LastName   string
	Age        int
	Occupation *Occupation
}
