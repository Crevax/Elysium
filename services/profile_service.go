package services

import "cjdavis.me/elysium/models"

type IProfileService interface {
	GetProfile() models.Profile
}

type ProfileService struct{}

func NewProfileService() *ProfileService {
	return &ProfileService{}
}

func (s *ProfileService) GetProfile() *models.Profile {
	return &models.Profile{
		FirstName: "CJ",
		LastName:  "Davis",
		Age:       26,
		Occupation: &models.Occupation{
			CompanyName: "CQL, Inc.",
			Position:    "Software Developer",
		},
	}
}
