package main

import (
	"fmt"
	"todo-api/config" // Import your new config package
	"todo-api/internal/controller"
	"todo-api/internal/entity"
	"todo-api/internal/repository"
	"todo-api/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 1. Load configuration from .env
	config.LoadConfig()

	// 2. Setup Postgres using Environment Variables
	// We use os.Getenv to pull the values you defined in your .env file
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBConfig.Host,
		config.DBConfig.User,
		config.DBConfig.Pass,
		config.DBConfig.Name,
		config.DBConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to Postgres")
	}

	// 3. Run Migrations (Adds 'todos' table alongside 'channel_configs')
	db.AutoMigrate(&entity.Todo{})

	// 4. Initialize Redis Connection
	// This uses the singleton and the RedisConfig struct you just loaded
	rdb := config.GetRedisConnection()

	// 5. Dependency Injection
	todoRepo := repository.NewTodoRepository(db, rdb)
	todoSvc := service.NewTodoService(todoRepo)
	todoCtrl := controller.NewTodoController(todoSvc)

	// 6. Echo Setup
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	api := e.Group("/api/v1")
	{
		api.POST("/todos", todoCtrl.Create)
		api.GET("/todos", todoCtrl.GetAll)
		api.PUT("/todos/:id", todoCtrl.Update)
		api.DELETE("/todos/:id", todoCtrl.Delete)
	}

	// Start Server
	log.Info().Msg("Server is running on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
