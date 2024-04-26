package repository

import (
	"goRepositoryPattern/database/models"
	database "goRepositoryPattern/database/sqlc"
	"goRepositoryPattern/messages"
	"goRepositoryPattern/util"
	"goRepositoryPattern/validators"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthRepository interface {
	Register(ctx *gin.Context, arg validators.RegisterInput) (models.RegisterResponse, error)
	Login(ctx *gin.Context, arg validators.LoginInput) (models.LoginResponse, error)
}

func (r *Repository) Register(ctx *gin.Context, arg validators.RegisterInput) (models.RegisterResponse, error) {

	// check if user account exists (I think I should do it here - handle the business logic or repository - this is only supposed to create account)
	dbUser, _ := r.DB.GetUserByEmail(ctx, arg.Email)

	if dbUser.ID != 0 {
		return models.RegisterResponse{}, messages.ErrUserExists
	}

	// hash password
	hashedPassword, err := util.HashPassword(arg.Password)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	args := database.CreateUserParams{
		Firstname:      arg.FirstName,
		Lastname:       arg.LastName,
		Email:          arg.Email,
		HashedPassword: hashedPassword,
	}

	user, err := r.DB.CreateUser(ctx, args)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	response := models.RegisterResponse{
		ID:        user.ID,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (r *Repository) Login(ctx *gin.Context, arg validators.LoginInput) (models.LoginResponse, error) {
	user, err := r.DB.GetUserByEmail(ctx, arg.Email)
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
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}
