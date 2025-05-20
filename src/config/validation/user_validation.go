package validation

import (
	"encoding/json"
	"errors"

	errors_api "github.com/Railssa1/crud-go/src/config/errors"
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
func ValidateUserError(err error) *errors_api.ApiErrors {
	// Valida se o campo enviado é o tipo correto
	var jsonErr *json.UnmarshalTypeError

	// Valida se o campo está seguindo o padrão required, email etc
	var jsonValidationError validator.ValidationErrors

	if errors.As(err, &jsonErr) {
		return errors_api.NewBadRequestError("Invalid field type")
	} else if errors.As(err, &jsonValidationError) {
		errorCauses := []errors_api.Causes{}

		for _, e := range err.(validator.ValidationErrors) {
			cause := errors_api.Causes{
				Message: e.Translate(translator),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return errors_api.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	} else {
		return errors_api.NewBadRequestError("Error trying to convert fields")
	}
}
