package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/limiter"

	"unemployed/internal/unemploy"
	"unemployed/internal/validator"
)

func main() {
	validator := validator.NewStructValidator()

	app := fiber.New(fiber.Config{
		AppName:         "unemployed",
		StructValidator: validator,
		ServerHeader:    "",
	})

	// Security headers
	app.Use(helmet.New())

	// Rate limiting: 60 requests per minute per IP
	app.Use(limiter.New(limiter.Config{
		Max:        60,
		Expiration: 1 * time.Minute,
	}))

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET"},
	}))

	unemploy.SetupRoutes(app)

	go func() {
		if err := app.Listen(":8000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()
}
