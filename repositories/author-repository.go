package repositories

import (
	"log"

	"cjdavis.me/elysium/db"
	"cjdavis.me/elysium/models"
)

type AuthorRepository struct {
}

func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{}
}

func (self *AuthorRepository) GetAllAuthors() ([]models.Author, error) {
	authors := []models.Author{}

	rows, err := db.AppDB().Query(`
SELECT id, first_name, last_name
FROM author
`)
	if err != nil {
		return authors, err
	}
	defer rows.Close()

	for rows.Next() {
		author := models.Author{}
		err := rows.Scan(&author.ID, &author.FirstName, &author.LastName)
		if err != nil {
			log.Println("Error scanning authors: ", err)
			continue
		}

		authors = append(authors, author)
	}

	err = rows.Err()
	if err != nil {
		log.Println("Error scanning authors: ", err)
	}

	return authors, err
}
