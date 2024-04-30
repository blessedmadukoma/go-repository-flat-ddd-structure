package repository

import (
	"goRepositoryPattern/database/models"
	"goRepositoryPattern/messages"
	"goRepositoryPattern/util"
	"goRepositoryPattern/validators"
	"time"

	"github.com/gin-gonic/gin"
)

func (r Repository) AdminLogin(ctx *gin.Context, arg validators.AdminLoginInput) (models.LoginResponse, error) {
	user, err := r.DB.GetAccountByEmail(ctx, arg.Email)
	if err != nil {
		return models.LoginResponse{}, messages.ErrUserNotExists
	}

	// check if user account exists
	if user.ID < 1 {
		return models.LoginResponse{}, messages.ErrInvalidCredentials
	}

	// check hashed password
	err = util.CheckPassword(arg.Password, user.HashedPassword)
	if err != nil {
		return models.LoginResponse{}, messages.ErrInvalidPassword
	}

	token, err := r.Token.CreateToken(user.ID, time.Minute*15)
	if err != nil {
		return models.LoginResponse{}, err
	}

	response := models.LoginResponse{
		ID:        user.ID,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
		Role:      user.Role,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}
