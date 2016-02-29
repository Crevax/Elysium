package services

import (
	"cjdavis.me/elysium/db"
	"cjdavis.me/elysium/models"
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
	tmpSession := db.AppDB().Copy()
	defer tmpSession.Close()

	profiles := tmpSession.DB(db.Database).C("profiles")
	var profile models.Profile
	profiles.Find(nil).One(&profile)
	return &profile
}
