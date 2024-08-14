package middlewares

import (
	"fmt"
	"strings"

	"banana-account-book.com/internal/config"
	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthHandler() fiber.Handler {
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
			c.Locals("userId", claims["userId"])
			return c.Next()
		}

		return c.Next()
	}
}
