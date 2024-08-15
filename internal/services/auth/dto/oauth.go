package dto

import "time"

type OauthRequestBody struct {
	Code string `json:"code" validate:"required"`
}

type OauthResponse struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
	Sync        bool      `json:"sync"`
}
