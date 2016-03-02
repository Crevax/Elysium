package services

import (
	"log"

	"cjdavis.me/elysium/models"
	"cjdavis.me/elysium/repositories"
)

type IProfileService interface {
	GetProfile() *models.Profile
}

type ProfileService struct {
}

func NewProfileService() *ProfileService {
	return &ProfileService{}
}

func (s *ProfileService) GetProfile() *models.Profile {
	profile, err := repositories.GetProfileRepository().GetProfile()
	if err != nil {
		log.Println(err)
	}

	return profile
}
