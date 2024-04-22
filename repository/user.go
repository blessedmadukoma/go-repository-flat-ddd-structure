package repository

import "github.com/gin-gonic/gin"

type UserRepository interface {
	GetUserByID(ctx *gin.Context, id string) (string, error)
}

func (r *Repository) GetUserByID(ctx *gin.Context, id string) (string, error) {
	user := &User{
		ID:       id,
		Name:     "Blessed",
		Email:    "blessed@gmail.com",
		Password: "password",
	}

	return user.Email, nil
}
