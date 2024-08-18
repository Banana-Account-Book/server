package validate

import (
	"fmt"
	"strings"

	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func ValidateDto(dto any) error {
	if err := validate.Struct(dto); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMsgs []string
			for _, err := range validationErrors {
				var errorMsg string
				switch err.Tag() {
				case "required":
					errorMsg = fmt.Sprintf("%s는(은) 필수 값입니다.", err.Field())
				case "email":
					errorMsg = fmt.Sprintf("%s는(은) 반드시 이메일 형식이어야 합니다.", err.Field())
				case "gte":
					errorMsg = fmt.Sprintf("%s는(은) 반드시 %s보다 크거나 같아야 합니다.", err.Field(), err.Param())
				case "lte":
					errorMsg = fmt.Sprintf("%s는(은) 반드시 %s보다 작거나 같아야 합니다.", err.Field(), err.Param())
				case "oneof":
					words := strings.Fields(err.Param())
					for i, word := range words {
						words[i] = "'" + word + "'"
					}
					errorMsg = fmt.Sprintf("%s는(은) 반드시 %s 중 하나여야 합니다.", err.Field(), strings.Join(words, ", "))
				default:
					errorMsg = fmt.Sprintf("%s는(은) 유효하지 않은 값입니다.", err.Field())
				}
				errorMsgs = append(errorMsgs, errorMsg)
			}
			message := strings.Join(errorMsgs, "\n")

			return appError.New(httpCode.BadRequest, message, message)
		}
	}
	return nil
}
