package services

import (
	"context"
	"github.com/Kartik30R/Tiket.git/models"
)

type AuthService struct {
	repo models.AuthRepository
}

// Login implements models.AuthService.
func (a *AuthService) Login(ctx context.Context, loginData *models.AuthCredentials) (string, *models.User, error) {
	panic("unimplemented")
}

// Register implements models.AuthService.
func (a *AuthService) Register(ctx context.Context, registerData *models.AuthCredentials) (string, *models.User, error) {
	panic("unimplemented")
}

func NewAuthServices(repo models.AuthRepository) models.AuthService {
	return &AuthService{
		repo: repo,
	}

}
