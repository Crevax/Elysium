package services

import (
	"log"

	"cjdavis.me/elysium/library"
	"cjdavis.me/elysium/repositories"
)

type AuthorService struct {
}

func NewAuthorService() *AuthorService {
	return &AuthorService{}
}

func (self *AuthorService) GetAllAuthors() []library.Author {
	authors, err := repositories.GetAuthorRepository().GetAllAuthors()
	if err != nil {
		log.Println(err)
	}

	return authors
}
