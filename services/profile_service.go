package services

import (
	"log"

	"cjdavis.me/elysium/profile"
	"cjdavis.me/elysium/repositories"
)

type ProfileService struct {
}

func NewProfileService() *ProfileService {
	return &ProfileService{}
}

func (s *ProfileService) GetProfile() *profile.Profile {
	profile, err := repositories.GetProfileRepository().GetProfile()
	if err != nil {
		log.Println(err)
	}

	return profile
}
