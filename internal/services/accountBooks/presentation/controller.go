package presentation

import (
	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/libs/validate"
	"banana-account-book.com/internal/services/accountBooks/application"
	"banana-account-book.com/internal/services/accountBooks/dto"
	"banana-account-book.com/internal/services/users/domain"
	"github.com/gofiber/fiber/v2"
)

type AccountBookController struct {
	accountBookService *application.AccountBookService
}

func NewAccountBookController(accountBookService *application.AccountBookService) *AccountBookController {
	return &AccountBookController{
		accountBookService: accountBookService,
	}
}

func (c *AccountBookController) Route(r fiber.Router) {
	r.Post("/", c.add)
}

func (c *AccountBookController) add(ctx *fiber.Ctx) error {
	// 1. ctx destructuring
	var body dto.AddAccountBookRequestBody
	user := ctx.Locals("user").(*domain.User)

	if err := ctx.BodyParser(&body); err != nil {
		return appError.Wrap(err)
	}

	if err := validate.ValidateDto(body); err != nil {
		return appError.Wrap(err)
	}

	// 2. call application service method
	if err := c.accountBookService.Add(user.Id, body.Name); err != nil {
		return appError.Wrap(err)
	}

	return ctx.Status(httpCode.Created.Code).SendString("created")
}
