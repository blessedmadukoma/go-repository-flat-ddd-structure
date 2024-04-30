package services

import (
	"goRepositoryPattern/database/models"
	"goRepositoryPattern/validators"

	"github.com/gin-gonic/gin"
)

func (s *AuthService) AdminLogin(ctx *gin.Context, arg validators.AdminLoginInput) (models.LoginResponse, error) {

	if err := arg.Validate(); err != nil {
		return models.LoginResponse{}, err
	}

	userResponse, err := s.repo.AdminLogin(ctx, arg)
	if err != nil {
		return models.LoginResponse{}, err
	}

	return userResponse, nil
}
