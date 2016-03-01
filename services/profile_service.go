package services

import (
	"database/sql"
	"encoding/json"
	"log"

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
	profile := models.Profile{}
	jsonRow := json.RawMessage{}

	err := db.AppDB().QueryRow(`
SELECT data FROM profile
ORDER BY data->>'id' ASC
LIMIT 1
`).Scan((*[]byte)(&jsonRow))
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("Error retrieving profile: " + err.Error())
			return &profile
		}
	}

	err = json.Unmarshal(jsonRow, &profile)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("Error unmarshalling json: " + err.Error())
			return &profile
		}
	}

	return &profile
}
