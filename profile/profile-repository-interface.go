package profile

type IProfileRepository interface {
	GetProfile() (*Profile, error)
}
