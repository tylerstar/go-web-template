package models

type (
	GetUserRequest struct {}

	GetUserResponse struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
)
