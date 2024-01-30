package internalerrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	valdiate := validator.New()
	err := valdiate.Struct(obj)
	if err == nil {
		return nil
	}
	validationErrors := err.(validator.ValidationErrors)
	valdationErro := validationErrors[0]

	field := strings.ToLower(valdationErro.StructField())

	switch valdationErro.Tag() {
	case "required":
		return errors.New(field + " is required")

	case "max":
		return errors.New(field + " is required with max " + valdationErro.Param())

	case "min":
		return errors.New(field + " is required with min " + valdationErro.Param())

	case "email":
		return errors.New(field + " is invalid")
	}
	return nil
}
