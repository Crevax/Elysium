package services

import "cjdavis.me/elysium/interfaces/services"

var profileService interfaces.IProfileService

func init() {
	profileService = NewProfileService()
}

func GetProfileService() interfaces.IProfileService {
	return profileService
}
