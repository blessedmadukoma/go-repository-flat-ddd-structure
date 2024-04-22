package services

import (
	"errors"
	"goRepositoryPattern/database/models"
	"goRepositoryPattern/repository"
	"goRepositoryPattern/validators"
	"log"

	"github.com/gin-gonic/gin"
)

// type UserResponse struct {
// 	ID        int64     `json:"id"`
// 	Email     string    `json:"email"`
// 	Token     string    `json:"token"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

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

	log.Println("in auth.service.register")
	// check if user account exists

	// call register repository
	user, err := s.repo.Register(ctx, arg)

	if err != nil {
		log.Println("Error in authrepo.register", err)
		return models.UserResponse{}, err
	}

	return user, nil
}

// func (s *AuthService) Login(ctx *gin.Context, user db.AuthModelOrSomethingLikeThis) (UserResponse, error) {
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
