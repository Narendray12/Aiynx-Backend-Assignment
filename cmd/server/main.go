package main

import (
	"database/sql"
	"log"

	"aiynx/config"
	"aiynx/internal/handler"
	"aiynx/internal/logger"
	"aiynx/internal/middleware"
	"aiynx/internal/repository"
	"aiynx/internal/routes"
	"aiynx/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	logger.Init()

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established")
	if err := db.Ping(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	app := fiber.New()
	app.Use(recover.New())
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	routes.Register(app, h)

	log.Printf("Server running on port %s\n", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
