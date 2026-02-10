package main

import (
    "net/http"
    "todo-api/config"
    "todo-api/internal/controller"
    "todo-api/internal/repository"
    "todo-api/internal/service"

    "github.com/go-playground/validator/v10"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/rs/zerolog/log"
)

type RequestValidator struct {
    validator *validator.Validate
}

func (rv *RequestValidator) Validate(i interface{}) error {
    if err := rv.validator.Struct(i); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return nil
}

func main() {
    config.LoadConfig()
    
    // Initialize DB & Redis (Assuming these return your GORM/Redis clients)
    var db = config.InitConnectDB(
        config.PostgreSQLConfig.PostgreSQLHost, 
        config.PostgreSQLConfig.PostgreSQLUser, 
        config.PostgreSQLConfig.PostgreSQLPassword, 
        config.PostgreSQLConfig.PostgreSQLDBName, 
        config.PostgreSQLConfig.PostgreSQLDBSchema, 
        config.PostgreSQLConfig.PostgreSQLPort,
    )

    if db == nil {
        log.Fatal().Msg("Database connection is null")
    }

    rdb := config.GetRedisConnection()

	
    e := echo.New()
    
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
	
    e.Validator = &RequestValidator{validator: validator.New()}
	
    todoRepo := repository.NewTodoRepository(db)
    todoSvc := service.NewTodoService(todoRepo, rdb)
    todoCtrl := controller.NewTodoController(todoSvc)

    todoCtrl.Routes(e)

    log.Info().Msg("Server is running on :8080")
    e.Logger.Fatal(e.Start(":8080"))
}