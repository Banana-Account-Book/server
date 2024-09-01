package presentation

import (
	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/libs/validate"
	"banana-account-book.com/internal/services/accountBooks/application"
	"banana-account-book.com/internal/services/accountBooks/dto"
	roleModel "banana-account-book.com/internal/services/roles/domain"
	userModel "banana-account-book.com/internal/services/users/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	r.Delete("/:id", c.delete)
}

// @Summary 가계부 생성
// @Description 로그인 한 사용자의 가계부를 생성한다.
// @Tags accountBooks
// @Accept json
// @Produce plain
// @Param body body dto.AddAccountBookRequestBody true "Account Book details"
// @Success 201 {string} string "created"
// @Failure 400 {object} appError.ErrorResponse "Bad Request"
// @Failure 500 {object} appError.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /account-books [post]
func (c *AccountBookController) add(ctx *fiber.Ctx) error {
	// 1. ctx destructuring
	var body dto.AddAccountBookRequestBody
	user := ctx.Locals("user").(*userModel.User)

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

// @Summary 가계부 삭제
// @Description 로그인 한 사용자의 가계부를 삭제한다.
// @Tags accountBooks
// @Accept json
// @Produce json
// @Param id path string true "Account Book ID" format(uuid)
// @Success 200 {object} map[string]string "{"data": "success"}"
// @Failure 400 {object} appError.ErrorResponse "Bad Request"
// @Failure 401 {object} appError.ErrorResponse "Unauthorized"
// @Failure 403 {object} appError.ErrorResponse "Forbidden"
// @Failure 404 {object} appError.ErrorResponse "Not Found"
// @Failure 500 {object} appError.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /account-books/{id} [delete]
func (c *AccountBookController) delete(ctx *fiber.Ctx) error {
	// 1. ctx destructuring
	roles := ctx.Locals("roles").([]*roleModel.Role)
	accountBookId := ctx.Params("id")

	// 2. call application service method
	if err := c.accountBookService.Delete(roles, uuid.MustParse(accountBookId)); err != nil {
		return appError.Wrap(err)
	}

	return ctx.Status(httpCode.Ok.Code).JSON(fiber.Map{"data": "success"})
}
