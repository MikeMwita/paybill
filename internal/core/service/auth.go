package service

import (
	"github.com/MikeMwita/paybill/adapters"
)

type AuthService struct {
	userService adapters.UserService
}

func (a *AuthService) RegisterUser(username, email, password string) error {
	err := a.userService.RegisterUser(username, email, password)
	if err != nil {
		return err
	}
	// Handle successful registration
	return nil
}

func (a *AuthService) AuthenticateUser(username, password string) (bool, error) {
	authenticated, err := a.userService.AuthenticateUser(username, password)
	if err != nil {
		return false, err
	}
	// Handle successful authentication
	return authenticated, nil
}

func NewAuthService(userService adapters.UserService) adapters.UserService {
	return &AuthService{userService: userService}
}
