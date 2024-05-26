package models

type Request struct {
	Message string `json:"message" validate:"required,min=1"`
}
