package services

import (
	"log"

	"cjdavis.me/elysium/library"
	"cjdavis.me/elysium/repositories"
)

type BookService struct {
}

func NewBookService() *BookService {
	return &BookService{}
}

func (self *BookService) GetAllBooks() []library.Book {
	books, err := repositories.GetBookRepository().GetAllBooks()
	if err != nil {
		log.Println(err)
	}

	return books
}

func (self *BookService) GetBooksByAuthor(authorID int) []library.Book {
	books, err := repositories.GetBookRepository().GetBooksByAuthor(authorID)
	if err != nil {
		log.Println(err)
	}

	return books
}
