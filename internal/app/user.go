package app

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"management/internal/domain"
	"time"
)

func (a *App) GetUserByEmail(u *domain.User) (*domain.User, error) {
	u, err := a.repo.GetUserByEmail(u.Email)
	if err != nil {
		return nil, err
	}

	return &domain.User{Username: u.Username, Email: u.Email}, nil
}

func (a *App) GetUserByID(u *domain.User) (*domain.User, error) {
	u, err := a.repo.GetUserByID(u.ID)
	if err != nil {
		return nil, err
	}

	return &domain.User{Username: u.Username, Email: u.Email}, nil
}

func (a *App) CreateUser(u *domain.User) error {
	if len(u.Password) == 0 {
		return errors.New("password cannot be empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(hash)
	_, err = a.repo.CreateUser(u)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) GetUserToken(u *domain.User) (string, error) {
	user, err := a.repo.GetUserByEmail(u.Email)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(a.cfg.Auth.Exp)).Unix()

	t, err := token.SignedString([]byte(a.cfg.Auth.Secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (a *App) ValidateUserPassword(u *domain.User) (bool, error) {
	user, err := a.repo.GetUserByEmail(u.Email)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(u.Password))
	if err != nil {
		return false, nil
	}

	return true, nil
}