package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type (
	ValErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}
	XValidator struct {
		validator *validator.Validate
	}
)

func GetValidate() *validator.Validate {
	return validate
}

var validate = validator.New()

func (v XValidator) Validate(data interface{}) []ValErrorResponse {
	validationErrors := []ValErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ValErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true
			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func ValidationFormat(errs []ValErrorResponse) {
	errMsgs := make([]string, 0)

	for _, err := range errs {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"[%s]: '%v' | Needs to implement '%s'",
			err.FailedField,
			err.Value,
			err.Tag,
		))
	}
}
