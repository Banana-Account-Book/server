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

// SignUp godoc
// @Summary User sign up
// @Description Register a new user and return an access token
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.SignUpRequestBody true "Sign up information"
// @Success 201 {object} dto.SignUpResponse "Returns access token"
// @Failure 400 {object} appError.ErrorResponse "Bad request"
// @Failure 500 {object} appError.ErrorResponse "Internal server error"
// @Router /users/signup [post]
func (c *UserController) signUp(ctx *fiber.Ctx) error {
	// 1. ctx destructuring
	var body dto.SignUpRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return appError.New(httpCode.BadRequest, "Failed to parse request body", "Invalid Request body")
	}

	if err := appError.ValidateDto(body); err != nil {
		return appError.Wrap(err)
	}

	// 2. call application service method
	accessToken, err := c.userService.SignUp(body.Email, body.Password, body.Name)
	if err != nil {
		return appError.Wrap(err)
	}

	// 3. return response
	response := dto.SignUpResponse{AccessToken: accessToken}
	return ctx.Status(httpCode.Created.Code).JSON(response)
}
