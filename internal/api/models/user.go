package models

type (
	GetUserRequest struct {}

	GetUserResponse struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	CreateUserRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	GetTokenRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	GetTokenResponse struct {
		Token string `json:"token"`
	}
)
