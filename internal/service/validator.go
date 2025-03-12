package service

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type ModelValidator struct {
	validate *validator.Validate
}

func NewModelValidator() *ModelValidator {
	v := validator.New()
	err := v.RegisterValidation("regex_pattern", validatorRegexPatternValidation)
	if err != nil {
		return nil
	}
	return &ModelValidator{validate: v}
}

func (mv *ModelValidator) Validate(m interface{}) error {
	return mv.validate.Struct(m)
}

func validatorRegexPatternValidation(fl validator.FieldLevel) bool {
	param := fl.Param()
	field := fl.Field().String()
	matched, err := regexp.MatchString(param, field)
	if err != nil {
		return false
	}
	return matched
}
