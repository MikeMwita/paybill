package adapters

import (
	"context"
	"github.com/MikeMwita/paybill/internal/models"
)

type DBStorage interface {
	Register(ctx context.Context, username string, email string, password string) error
	Login(ctx context.Context, username string, password string) (string, error)
	Logout(ctx context.Context) error
	Refresh(ctx context.Context, refreshToken string) (string, error)
	RequestPasswordReset(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token string, newPassword string) error
	RequestEmailVerification(ctx context.Context, email string) error
	VerifyEmail(ctx context.Context, token string) error
	GetUserProfile(ctx context.Context) (models.UserProfile, error)
	UpdateUserProfile(ctx context.Context, profile models.UserProfile) error
}
