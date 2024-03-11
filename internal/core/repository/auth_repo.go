package repository

import (
	"context"
	"github.com/MikeMwita/paybill/adapters"
	"github.com/MikeMwita/paybill/internal/models"
)

type authRepo struct {
	dbStorage adapters.DBStorage
}

func (a authRepo) Register(ctx context.Context, username string, email string, password string) error {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) Login(ctx context.Context, username string, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) Logout(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) Refresh(ctx context.Context, refreshToken string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) RequestPasswordReset(ctx context.Context, email string) error {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) ResetPassword(ctx context.Context, token string, newPassword string) error {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) RequestEmailVerification(ctx context.Context, email string) error {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) VerifyEmail(ctx context.Context, token string) error {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) GetUserProfile(ctx context.Context) (models.UserProfile, error) {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) UpdateUserProfile(ctx context.Context, profile models.UserProfile) error {
	//TODO implement me
	panic("implement me")
}

func NewAuthRepo(dbStorage adapters.DBStorage) adapters.DBStorage {
	return &authRepo{dbStorage: dbStorage}
}
