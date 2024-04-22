package services

import (
	"goRepositoryPattern/repository"
	"log"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s *UserService) GetUserByID(ctx *gin.Context, id string) (string, error) {
	// business logic

	email, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		log.Fatal("error:", err)
		return "", err
	}

	return email, nil
}
