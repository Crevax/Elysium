package services

import (
	"cjdavis.me/elysium/library"
	"cjdavis.me/elysium/profile"
)

var profileService profile.IProfileService
var bookService library.IBookService
var authorService library.IAuthorService

func init() {
	profileService = NewProfileService()
	bookService = NewBookService()
	authorService = NewAuthorService()
}

func GetProfileService() profile.IProfileService {
	return profileService
}

func GetBookService() library.IBookService {
	return bookService
}

func GetAuthorService() library.IAuthorService {
	return authorService
}
