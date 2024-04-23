package services

import (
	"goRepositoryPattern/database/models"
	"goRepositoryPattern/repository"
	"goRepositoryPattern/validators"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(ctx *gin.Context, arg validators.RegisterInput) (models.RegisterResponse, error) {
	// validate user input
	if err := arg.Validate(); err != nil {
		return models.RegisterResponse{}, err
	}

	// call register repository
	user, err := s.repo.Register(ctx, arg)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	return user, nil
}

func (s *AuthService) Login(ctx *gin.Context, arg validators.LoginInput) (models.LoginResponse, error) {
	// Validate request parameters
	if err := arg.Validate(); err != nil {
		return models.LoginResponse{}, err
	}

	userResponse, err := s.repo.Login(ctx, arg)
	if err != nil {
		return models.LoginResponse{}, err
	}

	return userResponse, nil
}
