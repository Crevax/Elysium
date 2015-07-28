package services

import (
	"os"

	"cjdavis.me/elysium/models"
	"gopkg.in/mgo.v2"
)

type IProfileService interface {
	GetProfile() *models.Profile
}

type ProfileService struct {
	database *mgo.Session
}

func NewProfileService(d *mgo.Session) *ProfileService {
	return &ProfileService{d}
}

func (s *ProfileService) GetProfile() *models.Profile {
	profiles := s.database.DB(os.Getenv("MONGO_DB_NAME")).C("profiles")
	var profile models.Profile
	profiles.Find(nil).One(&profile)
	return &profile
}
