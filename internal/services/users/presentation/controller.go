package presentation

import (
	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/services/users/application"
	"banana-account-book.com/internal/services/users/dto"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *application.UserService
}

func NewUserController(userService *application.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) Route(r fiber.Router) {
	r.Post("/signup", c.signUp)
}

func (c *UserController) signUp(ctx *fiber.Ctx) error {
	// 1. ctx destructuring
	var dto dto.SignUpRequestBody
	if err := ctx.BodyParser(&dto); err != nil {
		return appError.New(httpCode.BadRequest, "Failed to parse request body", "Invalid Request body")
	}

	if err := appError.ValidateDto(dto); err != nil {
		return err
	}

	// 2. call application service method
	accessToken, err := c.userService.SignUp(dto.Email, dto.Password, dto.Name)
	if err != nil {
		return err
	}
	return ctx.Status(httpCode.Created.Code).JSON(fiber.Map{
		"accessToken": accessToken,
	})
}
