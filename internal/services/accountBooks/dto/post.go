package dto

type AddAccountBookRequestBody struct {
	Name string `json:"name" validate:"required"`
}
