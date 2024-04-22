package repository

import (
	"goRepositoryPattern/database/models"
	database "goRepositoryPattern/database/sqlc"
	"goRepositoryPattern/validators"
	"log"

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
	user := database.User{
		ID:             1,
		Email:          arg.Email,
		HashedPassword: arg.Password,
	}

	// user, err := r.DB.CreateUser(ctx, arg)
	// if err != nil {
	// 	return database.User{}, err
	// }

	// If credentials are valid, generate authentication token (you can use JWT, for example)
	authToken := generateAuthToken(arg.Email)

	response := models.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: authToken,
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
