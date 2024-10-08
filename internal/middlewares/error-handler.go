package middlewares

import (
	"fmt"

	appError "banana-account-book.com/internal/libs/app-error"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if e, ok := err.(*fiber.Error); ok {
		return ctx.Status(e.Code).JSON(fiber.Map{"message": e.Message})
	}

	e := appError.UnWrap(err)
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	//TODO: log error with something (e.g. Sentry, ELK, File, etc.)
	fmt.Println(e.Stack)

	errResponse := appError.ErrorResponse{Data: e.ClientMessage}

	return ctx.Status(e.Code).JSON(errResponse)
}
