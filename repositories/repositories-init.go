package repositories

import "cjdavis.me/elysium/interfaces/repositories"

var profileRepository interfaces.IProfileRepository
var bookRepository interfaces.IBookRepository
var authorRepository interfaces.IAuthorRepository

func init() {
	profileRepository = NewProfileRepository()
	bookRepository = NewBookRepository()
	authorRepository = NewAuthorRepository()
}

func GetProfileRepository() interfaces.IProfileRepository {
	return profileRepository
}

func GetBookRepository() interfaces.IBookRepository {
	return bookRepository
}

func GetAuthorRepository() interfaces.IAuthorRepository {
	return authorRepository
}
