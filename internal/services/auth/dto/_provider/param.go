package dto

type RequestParam struct {
	Provider string `json:"provider" validate:"oneof=google kakao naver"`
}
