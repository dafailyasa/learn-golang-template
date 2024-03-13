package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	Validate *validator.Validate
}

func NewValidator() *Validator {
	validate := validator.New()

	return &Validator{
		Validate: validate,
	}
}

var _ ValidatorApplication = (*Validator)(nil)

type ValidatorApplication interface {
	ValidateStruct(data interface{}) []string
}

func (v *Validator) ValidateStruct(data interface{}) []string {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(v.Validate, trans)

	msgErrs := []string{}

	if errs := v.Validate.Struct(data); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			msg := err.Translate(trans)
			msgErrs = append(msgErrs, msg)
		}
	}

	return msgErrs
}
