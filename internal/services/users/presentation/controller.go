package presentation

import (
	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/libs/validate"
	"banana-account-book.com/internal/services/users/application"
	"banana-account-book.com/internal/services/users/domain"
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
	r.Patch("/", c.update)
}

// user update godoc
// @Summary User Update
// @Description 유저 정보 수정
// @Tags users
// @Accept json
// @Produce json
// @Param name body dto.UpdateUserRequestBody "수정할 user name"
// @Success 200 {object} dto.UpdateUserResponse "Successfully updated"
// @Failure 400 {object} error "Bad request"
// @Failure 500 {object} error "Internal server error"
// @Router /users [patch]
func (c *UserController) update(ctx *fiber.Ctx) error {
	// 1. ctx destructuring
	user, ok := ctx.Locals("user").(*domain.User)
	if !ok {
		return appError.New(httpCode.Unauthorized, "Unauthorized", "")
	}

	// 2. dto validation
	var body dto.UpdateUserRequestBody

	if err := ctx.BodyParser(&body); err != nil {
		return appError.New(httpCode.BadRequest, "Invalid request body", "")
	}

	if err := validate.ValidateDto(body); err != nil {
		return appError.Wrap(err)
	}

	// 3. call application service method
	result, err := c.userService.Update(user.Id, body)
	if err != nil {
		return appError.Wrap(err)
	}

	return ctx.Status(httpCode.Ok.Code).JSON(result)
}
