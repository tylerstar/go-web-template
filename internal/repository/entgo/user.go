package entgo

import (
	"context"
	"management/internal/domain"
	"management/internal/repository/entgo/ent/user"
)

func (r *EntgoRepository) GetUserByEmail(email string) (*domain.User, error) {
	u, err := r.client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:	       u.ID,
		Username:  u.Username,
		Email:	   u.Email,
		Password:  u.PasswordHash,
	}, nil
}