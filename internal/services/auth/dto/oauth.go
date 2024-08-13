package dto

type OauthRequestBody struct {
	Code string `json:"code" validate:"required"`
}
