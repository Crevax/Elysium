package interfaces

import "cjdavis.me/elysium/models"

type IAuthorService interface {
	GetAllAuthors() []models.Author
}
