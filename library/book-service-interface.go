package library

type IBookService interface {
	GetAllBooks() []Book
	GetBooksByAuthor(authorID int) []Book
}
