package main

import (
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/wawancallahan/go-upload/internal/router"
)

type App struct {
	*fiber.App
}

func NewApp() *App {
	return &App{fiber.New()}
}

func main() {
	app := NewApp()

	app.Mount("/api", router.New())

	// Close any connections on interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		app.exit()
	}()

	// Start listening on the specified address
	err := app.Listen(":8000")
	if err != nil {
		app.exit()
	}
}
func (app *App) registerMiddlewares() {
	// Handle Panic
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New())
}

// Stop the Fiber application
func (app *App) exit() {
	_ = app.Shutdown()
}
