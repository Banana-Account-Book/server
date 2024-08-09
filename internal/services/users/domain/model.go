package domain

import (
	"time"

	appError "banana-account-book.com/internal/libs/app-error"
	"banana-account-book.com/internal/libs/entity"
	httpCode "banana-account-book.com/internal/libs/http/code"
	"banana-account-book.com/internal/libs/jwt"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	entity.SoftDeletableAggregate
	Id           uuid.UUID      `json:"id" gorm:"primaryKey; type:uuid"`
	Email        string         `json:"email" gorm:"unique;type:varchar(50)"`
	Name         string         `json:"name" gorm:"type:varchar(50)"`
	Providers    pq.StringArray `json:"providers" gorm:"type:text[];"`
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
		token, err := jwt.Sign(u.Id, time.Hour*24*7)
		accessTokenChan <- result{token: token, err: err}
	}()
	go func() {
		err := u.rotateRefreshToken()
		refreshTokenChan <- err
	}()

	accessTokenResult := <-accessTokenChan
	refreshTokenErr := <-refreshTokenChan

	if accessTokenResult.err != nil {
		return "", appError.Wrap(accessTokenResult.err)
	}

	if refreshTokenErr != nil {
		return "", appError.Wrap(refreshTokenErr)
	}

	return accessTokenResult.token, nil
}

func (u *User) rotateRefreshToken() error {
	refreshToken, err := jwt.Sign(nil, time.Hour*24*90)
	if err != nil {
		return appError.New(httpCode.InternalServerError, "Failed to rotate refresh token.", "")
	}
	u.RefreshToken = refreshToken
	return nil
}
