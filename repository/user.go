package repository

import "github.com/gin-gonic/gin"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
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
