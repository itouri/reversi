package domain

import "gopkg.in/go-playground/validator.v9"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}
