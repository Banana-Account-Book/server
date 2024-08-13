package validate

import (
	"fmt"

	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func ValidateDto(dto any) error {
	if err := validate.Struct(dto); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, err := range validationErrors {
				var errorMsg string
				switch err.Tag() {
				case "required":
					errorMsg = fmt.Sprintf("%s is required.", err.Field())
				case "email":
					errorMsg = fmt.Sprintf("%s must be a valid email address", err.Field())
				case "gte":
					errorMsg = fmt.Sprintf("%s must be greater than or equal to %s", err.Field(), err.Param())
				case "lte":
					errorMsg = fmt.Sprintf("%s must be less than or equal to %s", err.Field(), err.Param())
				default:
					errorMsg = fmt.Sprintf("%s is not valid", err.Field())
				}
				message += fmt.Sprintf("%s\n", errorMsg)
			}

			return appError.New(httpCode.BadRequest, message, message)
		}
	}
	return nil

}
