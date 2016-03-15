package interfaces

import "cjdavis.me/elysium/models"

type IAuthorRepository interface {
	GetAllAuthors() ([]models.Author, error)
}
