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
	// Verify the reset token is valid and find the associated user
	user, err := a.verifyResetToken(ctx, token)
	if err != nil {
		return err
	}

	// Validate the new password (e.g., check length, complexity)
	if err := validateNewPassword(newPassword); err != nil {
		return err
	}

	// Hash the new password before storing it
	hashedPassword, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	// Update the user's password in the database
	if err := a.updateUserPassword(ctx, user.ID, hashedPassword); err != nil {
		return err
	}

	// Invalidate the reset token so it can't be used again
	if err := a.invalidateResetToken(ctx, token); err != nil {
		return err
	}
	return nil

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
