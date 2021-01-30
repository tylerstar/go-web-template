package entgo

import (
	"context"
	"management/internal/domain"
	User "management/internal/repository/entgo/ent/user"
)

func (r *EntgoRepository) GetUserByEmail(email string) (*domain.User, error) {
	u, err := r.client.User.
		Query().
		Where(User.EmailEQ(email)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:	          u.ID,
		Username:  	  u.Username,
		Email:	      u.Email,
		PasswordHash: u.PasswordHash,
	}, nil
}

func (r *EntgoRepository) GetUserByID(id string) (*domain.User, error) {
	u, err := r.client.User.
		Query().
		Where(User.IDEQ(id)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:	          u.ID,
		Username:  	  u.Username,
		Email:	      u.Email,
		PasswordHash: u.PasswordHash,
	}, nil
}

func (r *EntgoRepository) CreateUser(u *domain.User) (*domain.User, error) {
	user, err := r.client.User.
		Create().
		SetEmail(u.Email).
		SetUsername(u.Username).
		SetPasswordHash(u.PasswordHash).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}, nil
}