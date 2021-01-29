package app

import (
	"management/internal/domain"
)

func (a *App) GetUser(u *domain.User) (*domain.User, error) {
	u, err := a.repo.GetUserByEmail(u.Email)
	if err != nil {
		return nil, err
	}

	return &domain.User{Username: u.Username, Email: u.Email}, nil
}