package validations

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateRequest(input any, v *validator.Validate) []error {
	var errorMessages []error

	if err := v.Struct(input); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, validationError := range validationErrors {
				switch validationError.Tag() {
				case "required":
					errorMessages = append(errorMessages, fmt.Errorf("campo %s é obrigatório", validationError.Field()))
				case "max":
					errorMessages = append(errorMessages, fmt.Errorf("campo %s excede o tamanho permitido de %s caracteres", validationError.Field(), validationError.Param()))
				case "min":
					errorMessages = append(errorMessages, fmt.Errorf("campo %s deve ter pelo menos %s caracteres", validationError.Field(), validationError.Param()))
				case "email":
					errorMessages = append(errorMessages, fmt.Errorf("campo %s deve ser um email válido", validationError.Field()))
				default:
					errorMessages = append(errorMessages, fmt.Errorf("erro no campo %s: %s", validationError.Field(), validationError.Tag()))
				}
			}
		} else {

			errorMessages = append(errorMessages, err)
		}
		return errorMessages
	}
	return nil
}
