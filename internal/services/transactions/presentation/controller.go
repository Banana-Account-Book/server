package presentation

import (
	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/libs/validate"
	"banana-account-book.com/internal/services/transactions/application"
	"banana-account-book.com/internal/services/transactions/dto"
	"banana-account-book.com/internal/services/users/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TransactionController struct {
	transactionService *application.TransactionService
}

func NewTransactionController(transactionService *application.TransactionService) *TransactionController {
	return &TransactionController{
		transactionService: transactionService,
	}
}

func (c *TransactionController) Route(r fiber.Router) {
	r.Post("/", c.add)
}

// @Summary 거래내역 생성
// @Description 거래내역을 생성한다.
// @Tags transactions
// @Accept json
// @Produce json
// @Param body body dto.CreateTransactionRequest true "Add Transaction information"
// @Param accountBookId path string true "account book id"
// @Success 200 {object} dto.UpdateUserResponse "Updated user information"
// @Failure 400 {object} appError.ErrorResponse "Bad Request"
// @Failure 401 {object} appError.ErrorResponse "Unauthorized"
// @Failure 500 {object} appError.ErrorResponse "Internal Server Error"
// @Security BearerAuth
// @Router /account-books/:accountBookId/transactions [post]
func (c *TransactionController) add(ctx *fiber.Ctx) error {
	// 1. ctx destructuring
	var body dto.CreateTransactionRequest
	user := ctx.Locals("user").(*domain.User)
	accountBookId := ctx.Params("accountBookId")

	if err := ctx.BodyParser(&body); err != nil {
		return appError.Wrap(err)
	}

	if err := validate.ValidateDto(body); err != nil {
		return appError.Wrap(err)
	}

	// 2. call application service method
	if err := c.transactionService.Add(user.Id, uuid.MustParse(accountBookId), body); err != nil {
		return appError.Wrap(err)
	}

	return ctx.Status(httpCode.Created.Code).SendString("created")
}
