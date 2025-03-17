package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/somphonee/go-fiber-hex/config"
	"github.com/somphonee/go-fiber-hex/internal/adapters/handlers"
	"github.com/somphonee/go-fiber-hex/internal/adapters/repositories/postgres"
	"github.com/somphonee/go-fiber-hex/internal/adapters/routes"
	"github.com/somphonee/go-fiber-hex/internal/core/services"
	"github.com/somphonee/go-fiber-hex/internal/infrastructure/database"
)



func main() {
	// โหลด Config
	cfg := config.LoadConfig()

	// เชื่อมต่อฐานข้อมูล
	db := database.NewPostgresDB(cfg)

	// สร้าง Repository
	userRepo := postgres.NewUserRepository(db)

	// สร้าง Service
	userService := services.NewUserService(userRepo)

	// สร้าง Handler
	userHandler := handlers.NewUserHandler(userService)

	// สร้าง Fiber App
	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	// ตั้งค่า Routes
	routes.SetupRoutes(app, userHandler)

	// เริ่มต้น Server
	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("Server is running on port %s", cfg.AppPort)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}