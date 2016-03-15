package interfaces

import "cjdavis.me/elysium/models"

type IBookRepository interface {
	GetAllBooks() ([]models.Book, error)
	GetBooksByAuthor(authorID int) ([]models.Book, error)
}
