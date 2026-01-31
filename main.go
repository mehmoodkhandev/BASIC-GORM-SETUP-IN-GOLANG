package main

import (
	"fmt"
	"log"

	"github.com/MehmoodNadeemKhan1/CRUD_API/config"
	"github.com/MehmoodNadeemKhan1/CRUD_API/handler"
	"github.com/MehmoodNadeemKhan1/CRUD_API/model"
	"github.com/MehmoodNadeemKhan1/CRUD_API/repository"
	"github.com/MehmoodNadeemKhan1/CRUD_API/router"
	"github.com/MehmoodNadeemKhan1/CRUD_API/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	fmt.Println("Run Services...")
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	fmt.Println("Config loaded successfully:", loadConfig)
	db := config.ConnectionDB(loadConfig)
	fmt.Println("Database connected successfully")
	validate := validator.New()
	fmt.Println("Validator initialized successfully")
	db.AutoMigrate(&model.AuthModel{})
	fmt.Println("Database migrated successfully")
	//Init Repository
	authRepository := repository.NewAuthRepositoryImpl(db)
	fmt.Println("Auth repository initialized successfully")
	//Init Services
	authService := services.NewAuthServiceImpl(authRepository, validate)
	fmt.Println("Auth service initialized successfully")
	//Init Handler
	authHandler := handler.NewAuthHandler(authService)
	fmt.Println("Auth handler initialized successfully")
	//Init Router
	authRouter := router.NewRouter(authHandler)
	fmt.Println("Auth router initialized successfully")
	app.Mount("/api", authRouter)
	//Start Server
	fmt.Println("Starting server on port 8000...")

	log.Fatal(app.Listen(":8000"))

}
