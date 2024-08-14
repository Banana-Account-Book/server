package presentation

import (
	"banana-account-book.com/internal/services/users/application"
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

}
