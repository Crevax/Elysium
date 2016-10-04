package repositories

import "cjdavis.me/elysium/interfaces/repositories"
import "cjdavis.me/elysium/library"

var profileRepository interfaces.IProfileRepository
var bookRepository library.IBookRepository
var authorRepository library.IAuthorRepository

func init() {
	profileRepository = NewProfileRepository()
	bookRepository = NewBookRepository()
	authorRepository = NewAuthorRepository()
}

func GetProfileRepository() interfaces.IProfileRepository {
	return profileRepository
}

func GetBookRepository() library.IBookRepository {
	return bookRepository
}

func GetAuthorRepository() library.IAuthorRepository {
	return authorRepository
}
