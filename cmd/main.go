package main

import (
	"github.com/Steelyphil1/Personal_Website_API/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3001")
}
