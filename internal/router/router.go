package router

import (
	_ "banana-account-book.com/docs"

	"banana-account-book.com/internal/libs/health"
	"banana-account-book.com/internal/middlewares"
	accountBookPresentation "banana-account-book.com/internal/services/accountBooks/presentation"
	authPresentation "banana-account-book.com/internal/services/auth/presentation"
	transactionPresentation "banana-account-book.com/internal/services/transactions/presentation"
	userPresentation "banana-account-book.com/internal/services/users/presentation"
	"github.com/gofiber/fiber/v2"
)

// @title			Banana Account Book API
// @version		1.0
// @description	API Server for Banana Account Book
// @contact.name	API Support
// @contact.email	hch950627@naver.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath		/
func Route(
	r *fiber.App,
	userController *userPresentation.UserController,
	authController *authPresentation.AuthController,
	authHandler *middlewares.AuthHandler,
	accountBookController *accountBookPresentation.AccountBookController,
	transactionController *transactionPresentation.TransactionController,
) {
	health.Check(r)

	// public router
	publicRouter := r.Group("/")
	authRoute := publicRouter.Group("/auth")
	authController.Route(authRoute)

	// private router
	privateRouter := r.Group("/", authHandler.Auth())
	userRoute := privateRouter.Group("/users")
	userController.Route(userRoute)

	accountBookRoute := privateRouter.Group("/account-books")
	accountBookController.Route(accountBookRoute)

	transactionRoute := privateRouter.Group("/account-books/:accountBookId/transactions")
	transactionController.Route(transactionRoute)
}
