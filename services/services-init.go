package services

import (
	"cjdavis.me/elysium/interfaces/services"
	"cjdavis.me/elysium/library"
)

var profileService interfaces.IProfileService
var bookService library.IBookService
var authorService library.IAuthorService

func init() {
	profileService = NewProfileService()
	bookService = NewBookService()
	authorService = NewAuthorService()
}

func GetProfileService() interfaces.IProfileService {
	return profileService
}

func GetBookService() library.IBookService {
	return bookService
}

func GetAuthorService() library.IAuthorService {
	return authorService
}
