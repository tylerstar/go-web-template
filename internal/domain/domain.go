package domain

type (
	User struct {
		ID string
		Service bool
		Username string
		Password string
		PasswordHash string
		Email string
	}
)
