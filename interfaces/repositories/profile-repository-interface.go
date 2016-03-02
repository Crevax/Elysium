package interfaces

import "cjdavis.me/elysium/models"

type IProfileRepository interface {
	GetProfile() (*models.Profile, error)
}
