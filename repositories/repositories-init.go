package repositories

import (
	"cjdavis.me/elysium/library"
	"cjdavis.me/elysium/profile"
)

var profileRepository profile.IProfileRepository
var bookRepository library.IBookRepository
var authorRepository library.IAuthorRepository

func init() {
	profileRepository = NewProfileRepository()
	bookRepository = NewBookRepository()
	authorRepository = NewAuthorRepository()
}

func GetProfileRepository() profile.IProfileRepository {
	return profileRepository
}

func GetBookRepository() library.IBookRepository {
	return bookRepository
}

func GetAuthorRepository() library.IAuthorRepository {
	return authorRepository
}
