package library

type IBookRepository interface {
	GetAllBooks() ([]Book, error)
	GetBooksByAuthor(authorID int) ([]Book, error)
}
