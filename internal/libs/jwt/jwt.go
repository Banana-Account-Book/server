package jwt

import (
	"fmt"
	"time"

	"banana-account-book.com/internal/config"
	appError "banana-account-book.com/internal/libs/app-error"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"github.com/golang-jwt/jwt"
)

func Sign(userId any, expiredAfter time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", appError.New(httpCode.InternalServerError, "Failed to encode access token. Can not convert claims to jwt.MapClaims.", "")
	}

	if userId != nil {
		claims["userId"] = userId
	}
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(expiredAfter).Unix()

	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", appError.New(httpCode.InternalServerError, fmt.Sprintf("Failed to encode access token. %v", err), "")
	}

	return tokenString, nil
}

var AccessTokenExpiredAfter = time.Hour * 24
var RefreshTokenExpiredAfter = time.Hour * 24 * 90
