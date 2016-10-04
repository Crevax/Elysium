package library

type IAuthorRepository interface {
	GetAllAuthors() ([]Author, error)
}
