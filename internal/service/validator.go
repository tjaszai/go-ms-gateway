package service

import (
	baseValidator "github.com/go-playground/validator/v10"
	"regexp"
)

type Validator struct {
	validate *baseValidator.Validate
}

func NewValidator() *Validator {
	v := baseValidator.New()
	err := v.RegisterValidation("regex_pattern", validatorRegexPatternValidation)
	if err != nil {
		return nil
	}
	return &Validator{validate: v}
}

func (mv *Validator) ValidateObject(o interface{}) error {
	return mv.validate.Struct(o)
}

func validatorRegexPatternValidation(fl baseValidator.FieldLevel) bool {
	param := fl.Param()
	field := fl.Field().String()
	matched, err := regexp.MatchString(param, field)
	if err != nil {
		return false
	}
	return matched
}
