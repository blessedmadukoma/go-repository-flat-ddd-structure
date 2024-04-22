package repository

import (
	"goRepositoryPattern/database/models"
	database "goRepositoryPattern/database/sqlc"
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
	// Register(ctx *gin.Context, email, password string) (*User, error)
	// Login(ctx *gin.Context, email, password string) (*User, error)
	// Register(ctx *gin.Context, arg database.CreateUserParams) (database.User, error)
	Register(ctx *gin.Context, arg validators.RegisterInput) (models.UserResponse, error)
	// Login(ctx *gin.Context, email, password string) (database.User, error)
	Login(ctx *gin.Context, email, password string) (string, error)
}

func (r *Repository) Register(ctx *gin.Context, arg validators.RegisterInput) (models.UserResponse, error) {
	// user := database.User{
	// 	ID:             1,
	// 	Email:          arg.Email,
	// 	HashedPassword: arg.Password,
	// }
	log.Println("in auth.repository.register")

	// hash password
	hashedPassword, err := util.HashPassword(arg.Password)
	if err != nil {
		log.Fatal("unable to hash password:", err)
		return models.UserResponse{}, err
	}

	args := database.CreateUserParams{
		Firstname:      arg.FirstName,
		Lastname:       arg.LastName,
		Email:          arg.Email,
		HashedPassword: hashedPassword,
	}

	log.Println("creating user")
	user, err := r.DB.CreateUser(ctx, args)
	if err != nil {
		log.Println("unable to create user:", err)
		return models.UserResponse{}, err
	}

	log.Println("generating token")
	// If credentials are valid, generate authentication token (you can use JWT, for example)
	token, err := r.Token.CreateToken(user.ID, time.Minute*15)
	// token, err := r.Token.CreateToken(user.ID)
	if err != nil {
		log.Println("unable to create token:", err)
		return models.UserResponse{}, err
	}

	response := models.UserResponse{
		ID:        user.ID,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
		Token:     token,
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
