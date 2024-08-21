package middlewares

import (
	"fmt"
	"strings"

	"banana-account-book.com/internal/config"
	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	roleModel "banana-account-book.com/internal/services/roles/domain"
	roleInfra "banana-account-book.com/internal/services/roles/infrastructure"
	userModel "banana-account-book.com/internal/services/users/domain"
	userInfra "banana-account-book.com/internal/services/users/infrastructure"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type AuthHandler struct {
	userRepository userInfra.UserRepository
	roleRepository roleInfra.RoleRepository
}

func NewAuthHandler(userRepository userInfra.UserRepository, roleRepository roleInfra.RoleRepository) *AuthHandler {
	return &AuthHandler{
		userRepository: userRepository,
		roleRepository: roleRepository,
	}
}

func (a *AuthHandler) Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		splitToken := strings.Split(auth, " ")
		if splitToken[0] != "Bearer" {
			return appError.New(httpCode.Unauthorized, "Invalid token type", "")
		}
		accessToken := splitToken[1]

		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, appError.New(httpCode.Unauthorized, fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]), "")
			}
			return []byte(config.SecretKey), nil
		})

		if err != nil {
			return appError.New(httpCode.Unauthorized, "Invalid Token", "")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user, roles, err := a.getUserAndRoles(uuid.MustParse(claims["id"].(string)))

			if err != nil {
				return appError.Wrap(err)
			}

			c.Locals("user", user)
			c.Locals("roles", roles)

			return c.Next()
		}

		return c.Next()
	}
}

func (a *AuthHandler) getUserAndRoles(userId uuid.UUID) (*userModel.User, []*roleModel.Role, error) {
	user, err := a.userRepository.FindOneOrFail(nil, userId)
	if err != nil {
		return nil, nil, appError.New(httpCode.Unauthorized, fmt.Sprintf("Unauthorized: %v", err), "")
	}

	roles, _, err := a.roleRepository.FindByUserId(nil, userId)

	if err != nil {
		return nil, nil, appError.New(httpCode.Unauthorized, fmt.Sprintf("Unauthorized: %v", err), "")
	}

	return user, roles, nil
}
