package dtos

type UpdateUserPasswordCommandDTO struct {
	OldPassword string
	NewPassword string
	Email       string
	ID          string
}
