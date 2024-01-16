package dtos

type BasicLoginCommandDTO struct {
	Identifier string `validate:"required"`
	Password   string `validate:"required"`
}

type ChangePasswordCommandDTO struct {
	OldPassword string `validate:"required"`
	NewPassword string `validate:"required"`
}

type JWTPayload struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Role      string
}
