package repositories

import "cjdavis.me/elysium/interfaces/repositories"

var profileRepository interfaces.IProfileRepository

func init() {
	profileRepository = NewProfileRepository()
}

func GetProfileRepository() interfaces.IProfileRepository {
	return profileRepository
}
