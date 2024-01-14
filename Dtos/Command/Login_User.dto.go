package dtos

type BasicLoginCommandDTO struct {
	Identifier string `validate:"required"`
	Password   string `validate:"required"`
}
