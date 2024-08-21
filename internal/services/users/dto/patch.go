package dto

import "github.com/google/uuid"

type UpdateUserRequestBody struct {
	Name string `json:"name"`
}

type UpdateUserResponse struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}
