package models

type Book struct {
	ID     int
	Title  string
	Author Author
	Read   bool
}
