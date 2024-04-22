package repository

import (
	"goRepositoryPattern/database/models"
	database "goRepositoryPattern/database/sqlc"
	"goRepositoryPattern/messages"
	"goRepositoryPattern/util"
	"goRepositoryPattern/validators"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

// Auth
type AuthRepository interface {
	Register(ctx *gin.Context, arg validators.RegisterInput) (models.UserResponse, error)
	Login(ctx *gin.Context, email, password string) (string, error)
}

func (r *Repository) Register(ctx *gin.Context, arg validators.RegisterInput) (models.UserResponse, error) {

	// check if user account exists (I think I should do it here - handle the business logic or repository - this is only supposed to create account)
	dbUser, _ := r.DB.GetUserByEmail(ctx, arg.Email)

	if dbUser.ID != 0 {
		return models.UserResponse{}, messages.ErrUserExists
	}

	// hash password
	hashedPassword, err := util.HashPassword(arg.Password)
	if err != nil {
		return models.UserResponse{}, err
	}

	args := database.CreateUserParams{
		Firstname:      arg.FirstName,
		Lastname:       arg.LastName,
		Email:          arg.Email,
		HashedPassword: hashedPassword,
	}

	user, err := r.DB.CreateUser(ctx, args)
	if err != nil {
		return models.UserResponse{}, err
	}

	token, err := r.Token.CreateToken(user.ID, time.Minute*15)
	if err != nil {
		return models.UserResponse{}, err
	}

	response := models.UserResponse{
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

// func (r *Repository) Login(ctx *gin.Context, email, password string) (*User, error) {
func (r *Repository) Login(ctx *gin.Context, email, password string) (string, error) {
	// user := &User{
	// 	ID:       "2",
	// 	Name:     "Blessed 2 logged in",
	// 	Email:    email,
	// 	Password: password,
	// }

	user := database.User{
		ID:             1,
		Email:          email,
		HashedPassword: password,
	}

	// return user, nil
	return user.Email, nil
}

func generateAuthToken(username string) string {
	log.Println("got to auth token")
	// Implement token generation logic here
	return "your_generated_token" + username
}
