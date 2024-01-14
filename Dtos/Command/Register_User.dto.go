package dtos

type RegisterUserCommandDTO struct {
	ID                string
	MemberNumber      string // Uses sql.NullString to handle nullable strings
	Email             string
	FirstName         string
	LastName          string
	ProfilePictureUrl string
}
