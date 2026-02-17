package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v3"

	"unemployed/internal/unemploy"
	"unemployed/internal/validator"
)

func main() {
	validator := validator.NewStructValidator()
	engine := html.New("./internal/views", ".html")

	app := fiber.New(fiber.Config{
		AppName:         "unemployed",
		Views:           engine,
		ViewsLayout:     "layouts/main",
		StructValidator: validator,
	})

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
