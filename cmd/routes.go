package main

import (
	"github.com/Steelyphil1/Personal_Website_API/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
