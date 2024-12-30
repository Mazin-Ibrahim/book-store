package http

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func validationErrors(err error) ([]byte, error) {
	validationErrors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()
		validationErrors[field] = fmt.Sprintf("Validation failed on '%s' tag", tag)
	}
	errorJSON, err := json.Marshal(validationErrors)

	if err != nil {
		return nil, err
	}
	return errorJSON, nil
}
