package services

import (
	"log"

	"cjdavis.me/elysium/models"
	"cjdavis.me/elysium/repositories"
)

type AuthorService struct {
}

func NewAuthorService() *AuthorService {
	return &AuthorService{}
}

func (self *AuthorService) GetAllAuthors() []models.Author {
	authors, err := repositories.GetAuthorRepository().GetAllAuthors()
	if err != nil {
		log.Println(err)
	}

	return authors
}
