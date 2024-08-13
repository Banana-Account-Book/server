package presentation

import (
	"time"

	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/libs/jwt"
	"banana-account-book.com/internal/libs/validate"
	"banana-account-book.com/internal/services/auth/application"
	"banana-account-book.com/internal/services/auth/dto"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService *application.AuthService
}

func NewAuthController(authService *application.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Route(r fiber.Router) {
	r.Get("/:provider", c.GetLink)
	r.Post("/:provider", c.Callback)
}

func (c *AuthController) GetLink(ctx *fiber.Ctx) error {
	// 1. ctx destructuring
	provider := ctx.Params("provider")

	// 2. call application service method
	url, err := c.authService.GetAuthUrl(provider)
	if err != nil {
		return appError.Wrap(err)
	}

	// 3. return response
	return ctx.Status(httpCode.Ok.Code).JSON(fiber.Map{"url": url})

}

func (c *AuthController) Callback(ctx *fiber.Ctx) error {
	// 1. ctx destructuring
	provider := ctx.Params("provider")
	var body dto.OauthRequestBody

	// 2. parse request body
	if err := ctx.BodyParser(&body); err != nil {
		return appError.New(httpCode.BadRequest, "Invalid request body", "")
	}

	accessToken, err := c.authService.OAuth(body.Code, provider)
	if err != nil {
		return appError.Wrap(err)
	}

	if err := validate.ValidateDto(body); err != nil {
		return appError.Wrap(err)
	}

	return ctx.Status(httpCode.Ok.Code).JSON(fiber.Map{"accessToken": accessToken, "expiredAt": time.Now().Add(jwt.AccessTokenExpiredAfter)})
}
