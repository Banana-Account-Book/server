package domain

import (
	appError "banana-account-book.com/internal/libs/app-error"
	"banana-account-book.com/internal/libs/ddd"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/libs/jwt"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ddd.SoftDeletableAggregate
	Id           uuid.UUID      `json:"id" gorm:"primaryKey; type:uuid"`
	Email        string         `json:"email" gorm:"unique;type:varchar(50); not null;"`
	Name         string         `json:"name" gorm:"type:varchar(50); not null;"`
	Providers    pq.StringArray `json:"providers" gorm:"type:text[];not null;"`
	RefreshToken string         `json:"refreshToken" gorm:"column:refreshToken;type:varchar(255)"`
}

func (u *User) TableName() string {
	return "user"
}

func New(email, name string, providers []string) (*User, error) {
	uuId, err := uuid.NewV7()
	if err != nil {
		return nil, appError.New(httpCode.InternalServerError, "Failed to create new user. Can not generate uuid.", "")
	}

	user := &User{
		Id:        uuId,
		Email:     email,
		Name:      name,
		Providers: pq.StringArray(providers),
	}

	return user, nil
}

func (u *User) EncodeAccessToken() (string, error) {
	type result struct {
		token string
		err   error
	}

	accessTokenChan := make(chan result, 1)
	refreshTokenChan := make(chan error, 1)
	go func() {
		token, err := jwt.Sign(u.Id, jwt.AccessTokenExpiredAfter)
		accessTokenChan <- result{token: token, err: err}
	}()
	go func() {
		err := u.rotateRefreshToken()
		refreshTokenChan <- err
	}()

	accessTokenResult := <-accessTokenChan
	refreshTokenErr := <-refreshTokenChan

	if accessTokenResult.err != nil {
		return "", appError.New(httpCode.Unauthorized, accessTokenResult.err.Error(), "")
	}

	if refreshTokenErr != nil {
		return "", appError.New(httpCode.Unauthorized, refreshTokenErr.Error(), "")
	}

	return accessTokenResult.token, nil
}

func (u *User) rotateRefreshToken() error {
	refreshToken, err := jwt.Sign(nil, jwt.RefreshTokenExpiredAfter)
	if err != nil {
		return appError.New(httpCode.InternalServerError, "Failed to rotate refresh token.", "")
	}
	u.RefreshToken = refreshToken
	return nil
}

func (u *User) HasProvider(provider string) bool {
	for _, p := range u.Providers {
		if p == provider {
			return true
		}
	}
	return false
}

func (u *User) AddProvider(provider string) {
	u.Providers = append(u.Providers, provider)
}
