package utils

import (
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/en"
)

func validateDurationMin(fl validator.FieldLevel) bool {
	param := fl.Param()
	minDuration, err := time.ParseDuration(param)
	if err != nil {
		return false
	}

	val, ok := fl.Field().Interface().(time.Duration)
	if !ok {
		return false
	}

	if val < minDuration {
		return false
	}

	return true
}

func validateDurationMax(fl validator.FieldLevel) bool {
	param := fl.Param()
	maxDuration, err := time.ParseDuration(param)
	if err != nil {
		return false
	}

	val, ok := fl.Field().Interface().(time.Duration)
	if !ok {
		return false
	}

	if val > maxDuration {
		return false
	}

	return true
}

// NewValidator новый валидатор
func NewValidator() (*validator.Validate, ut.Translator) {
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	if err := validate.RegisterValidation("duration-min", validateDurationMin); err != nil {
		panic(err)
	}

	if err := validate.RegisterValidation("duration-max", validateDurationMax); err != nil {
		panic(err)
	}

	err := translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}

	return validate, trans
}
