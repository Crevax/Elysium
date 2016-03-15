package services

import "cjdavis.me/elysium/interfaces/services"

var profileService interfaces.IProfileService
var bookService interfaces.IBookService
var authorService interfaces.IAuthorService

func init() {
	profileService = NewProfileService()
	bookService = NewBookService()
	authorService = NewAuthorService()
}

func GetProfileService() interfaces.IProfileService {
	return profileService
}

func GetBookService() interfaces.IBookService {
	return bookService
}

func GetAuthorService() interfaces.IAuthorService {
	return authorService
}
