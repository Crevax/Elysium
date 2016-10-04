package repositories

import (
	"log"

	"cjdavis.me/elysium/db"
	"cjdavis.me/elysium/library"
)

type BookRepository struct {
}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (self *BookRepository) GetAllBooks() ([]library.Book, error) {
	books := []library.Book{}

	rows, err := db.AppDB().Query(`
SELECT b.id, b.title, b.read, a.id, a.first_name, a.last_name
FROM book b
LEFT JOIN author a on a.id = b.authorid
`)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		book := library.Book{}
		author := library.Author{}
		err := rows.Scan(&book.ID, &book.Title, &book.Read, &author.ID, &author.FirstName, &author.LastName)
		if err != nil {
			log.Println("Error scanning books: ", err)
			continue
		}

		book.Author = author
		books = append(books, book)
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error scanning books: ", err)
	}

	return books, err
}

func (self *BookRepository) GetBooksByAuthor(authorID int) ([]library.Book, error) {
	books := []library.Book{}

	rows, err := db.AppDB().Query(`
SELECT b.id, b.title, b.read, a.id, a.first_name, a.last_name
FROM book b
LEFT JOIN author a on a.ID = b.authorid
WHERE b.authorid = $1
`, authorID)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		book := library.Book{}
		author := library.Author{}
		err := rows.Scan(&book.ID, &book.Title, &book.Read, &author.ID, &author.FirstName, &author.LastName)
		if err != nil {
			log.Println("Error scanning books: ", err)
			continue
		}

		book.Author = author
		books = append(books, book)
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error scanning books: ", err)
	}

	return books, err
}
