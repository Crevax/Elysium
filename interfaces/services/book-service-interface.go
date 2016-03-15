package interfaces

import "cjdavis.me/elysium/models"

type IBookService interface {
	GetAllBooks() []models.Book
	GetBooksByAuthor(authorID int) []models.Book
}
