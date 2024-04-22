package services

import (
	"errors"
	"goRepositoryPattern/database/models"
	"goRepositoryPattern/repository"
	"goRepositoryPattern/validators"
	"log"

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

func (s *AuthService) Register(ctx *gin.Context, arg validators.RegisterInput) (models.UserResponse, error) {
	if arg.Email == "" || arg.Password == "" {
		return models.UserResponse{}, errors.New("email and password are required")
	}

	// check if user account exists (I think I should do it here - handle the business logic or repository - this is only supposed to create account)

	// call register repository
	user, err := s.repo.Register(ctx, arg)

	if err != nil {
		return models.UserResponse{}, err
	}

	return user, nil
}

func (s *AuthService) Login(ctx *gin.Context, email, password string) (string, error) {
	// Validate request parameters
	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	// Call repository method to verify credentials
	log.Println("going into auth repo login with:", email, password)
	_, err := s.repo.Login(ctx, email, password)
	if err != nil {
		log.Println("error inside auth repo login:", err)
		return "", err
	}

	// If credentials are valid, generate authentication token (you can use JWT, for example)
	// authToken := generateAuthToken(email)

	authToken := "stringauthtoken"

	return authToken, nil
}
