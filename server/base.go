package server

import (
	"log"
	"os"

	// "goRepositoryPattern/db"
	database "goRepositoryPattern/database/sqlc"
	"goRepositoryPattern/repository"
	"goRepositoryPattern/routes"
	"goRepositoryPattern/services"
	"goRepositoryPattern/token"
	"goRepositoryPattern/util"

	// "goRepositoryPattern/tasks"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var engine *gin.Engine
var tokenController *token.JWTToken

// Initialize initiates the app instance
func Initialize() error {

	config := util.LoadEnvConfig()
	log.Println("Config:", config)

	if err := loadEnv(); err != nil {
		return err
	}

	tokenController = token.NewJWTToken(&config)

	// database connection
	DB, err := database.ConnectDataBase(config)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
		return err
	}

	// Run migrations
	// if err := db.RunMigrations(); err != nil {
	// 	return err
	// }

	engine = gin.Default()

	repo := repository.NewRepository(DB, tokenController)
	authService := services.NewAuthService(repo)
	userService := services.NewUserService(repo)

	// services := services.NewService(repo)
	services := services.NewService(authService, userService)

	// go func() {
	// 	if err := tasks.StartWorker(tasks.NewTask(repo, s)); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// Register routes
	routes.RegisterRoutes(engine, repo, services)

	return nil

}

func loadEnv() error {
	if os.Getenv("APP_ENV") == "dev" {
		return godotenv.Load()
	}
	return nil
}

// Run the app
func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Println("Server starting on port:" + port)
	engine.Run(":" + port)
}
