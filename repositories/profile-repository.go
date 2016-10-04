package repositories

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"cjdavis.me/elysium/db"
	"cjdavis.me/elysium/profile"
)

type ProfileRepository struct {
}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{}
}

func (s *ProfileRepository) GetProfile() (*profile.Profile, error) {
	profile := profile.Profile{}
	jsonRow := json.RawMessage{}

	err := db.AppDB().QueryRow(`
SELECT data FROM profile
ORDER BY data->>'id' ASC
LIMIT 1
`).Scan((*[]byte)(&jsonRow))
	if err != nil {
		if err != sql.ErrNoRows {
			return &profile, errors.New(fmt.Sprintf("Error retrieving profile: %s", err.Error()))
		}
	}

	err = json.Unmarshal(jsonRow, &profile)
	if err != nil {
		if err != sql.ErrNoRows {
			return &profile, errors.New(fmt.Sprintf("Error unmarshalling json: %s", err.Error()))
		}
	}

	return &profile, nil
}
