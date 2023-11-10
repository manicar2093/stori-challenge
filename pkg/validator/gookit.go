package validator

import (
	"github.com/google/uuid"
	"github.com/gookit/validate"
)

type GooKitValidator struct {
	StopOnError bool
}

const (
	uuidNotValidError = "is not valid for UUID type"
)

func NewGooKitValidator() *GooKitValidator {
	return &GooKitValidator{
		StopOnError: false,
	}
}

func (c *GooKitValidator) ValidateStruct(toValidate interface{}) error {
	v := c.configuratedValidator(toValidate)
	if v.Validate() {
		return nil
	}
	return &ValidationError{Errors: v.Errors}
}

func (c *GooKitValidator) Validate(i interface{}) error {
	return c.ValidateStruct(i)
}

func (c *GooKitValidator) configuratedValidator(toValidate interface{}) *validate.Validation {
	v := validate.Struct(toValidate) //nolint:varnamelen
	v.StopOnError = c.StopOnError
	v.AddValidator("googleUUID", func(val interface{}) bool {
		return val.(uuid.UUID) != uuid.Nil
	})
	v.AddMessages(map[string]string{
		"uuid":       uuidNotValidError,
		"isUUID":     uuidNotValidError,
		"googleUUID": uuidNotValidError,
		"required":   "needs to be on request",
	})
	return v
}
