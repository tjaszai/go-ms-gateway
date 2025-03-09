package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type checker struct {
	validate *validator.Validate
}

func (c *checker) ValidateDto(s interface{}) error {
	return c.validate.Struct(s)
}

func regexPatternValidation(fl validator.FieldLevel) bool {
	param := fl.Param()
	field := fl.Field().String()

	matched, err := regexp.MatchString(param, field)
	if err != nil {
		return false
	}
	return matched
}

func createChecker() *checker {
	checker := checker{validate: validator.New()}

	err := checker.validate.RegisterValidation("regex_pattern", regexPatternValidation)
	if err != nil {
		return nil
	}

	return &checker
}

var Checker = createChecker()
