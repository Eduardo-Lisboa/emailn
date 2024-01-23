package internalerrors

import (
	"errors"

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

	switch valdationErro.Tag() {
	case "required":
		return errors.New(valdationErro.StructField() + "is required")

	case "max":
		return errors.New(valdationErro.StructField() + "is required with max " + valdationErro.Param())

	case "min":
		return errors.New(valdationErro.StructField() + "is required with min " + valdationErro.Param())

	case "eamil":
		return errors.New(valdationErro.StructField() + "is invalid")
	}
	return nil
}
