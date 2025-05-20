package validation

import (
	"encoding/json"
	"errors"

	config "github.com/Railssa1/crud-go/src/config/errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
)

var (
	ValidatorInstance = validator.New()
	translator        ut.Translator
)

// Função responsável por traduzir os erros para inglês
func init() {
	if ginValidatorEngine, ok := binding.Validator.Engine().(*validator.Validate); ok {
		englishLocale := en.New()
		universalTranslator := ut.New(englishLocale, englishLocale)

		translator, _ = universalTranslator.GetTranslator("en")
		enTranslation.RegisterDefaultTranslations(ginValidatorEngine, translator)
	}
}

// Função responsável por validar os erros
func ValidateUserError(err error) *config.ApiErrors {
	// Valida se o campo enviado é o tipo correto
	var jsonErr *json.UnmarshalTypeError

	// Valida se o campo está seguindo o padrão required, email etc
	var jsonValidationError validator.ValidationErrors

	if errors.As(err, &jsonErr) {
		return config.NewBadRequestError("Invalid field type")
	} else if errors.As(err, &jsonValidationError) {
		errorCauses := []config.Causes{}

		for _, e := range err.(validator.ValidationErrors) {
			cause := config.Causes{
				Message: e.Translate(translator),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return config.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	} else {
		return config.NewBadRequestError("Error trying to convert fields")
	}
}
