package interfaces

import "cjdavis.me/elysium/models"

type IProfileService interface {
	GetProfile() *models.Profile
}
