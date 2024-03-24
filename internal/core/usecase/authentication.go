package usecase

import (
	"context"
	"errors"
	"github.com/MikeMwita/paybill/adapters"
	"github.com/MikeMwita/paybill/internal/models"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidEmail    = errors.New("invalid email")
)

type AuthUseCase struct {
	authService adapters.AuthAdapter
}

func (a AuthUseCase) Register(ctx context.Context, profile *models.UserProfile) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) Login(ctx context.Context, username string, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) Logout(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) Refresh(ctx context.Context, refreshToken string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) RequestPasswordReset(ctx context.Context, email string) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) ResetPassword(ctx context.Context, token string, newPassword string) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) RequestEmailVerification(ctx context.Context, email string) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) VerifyEmail(ctx context.Context, token string) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) GetUserProfile(ctx context.Context) (models.UserProfile, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) UpdateUserProfile(ctx context.Context, profile models.UserProfile) error {
	//TODO implement me
	panic("implement me")
}

//func (a AuthUseCase) Register(username string, email string, password string) error {
//	//validate  the password
//	if !validation.ValidatePassword(password) {
//		return ErrInvalidPassword
//	}
//	//validate the email
//	if !validation.ValidateEmail(email) {
//		return ErrInvalidEmail
//	}
//	return nil
//}

func NewAuthUseCase(authService adapters.AuthAdapter) adapters.AuthAdapter {
	return &AuthUseCase{authService: authService}
}
