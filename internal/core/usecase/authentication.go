package usecase

import (
	"errors"
	"github.com/MikeMwita/paybill/adapters"
	"github.com/MikeMwita/paybill/pkg/validation"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidEmail    = errors.New("invalid email")
)

type AuthUseCase struct {
	authService adapters.AuthAdapter
}

func (a AuthUseCase) Register(username string, email string, password string) error {
	//validate  the password
	if !validation.ValidatePassword(password) {
		return ErrInvalidPassword
	}
	//validate the email
	if !validation.ValidateEmail(email) {
		return ErrInvalidEmail
	}
	return nil
}

func (a AuthUseCase) Login(username string, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthUseCase(authService adapters.AuthAdapter) adapters.AuthAdapter {
	return &AuthUseCase{authService: authService}
}
