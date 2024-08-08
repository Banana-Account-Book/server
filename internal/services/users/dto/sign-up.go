package dto

type SignUpRequestBody struct {
	Email    string `json:"email" validate:"required,email" example:"hch950627@naver.com"`
	Password string `json:"password" validate:"required" example:"q1w2e3r4!@"`
	Name     string `json:"name" validate:"required" example:"arthur"`
}

type SignUpResponse struct {
	AccessToken string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
