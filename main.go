package main

import (
	"github.com/basit9958/shorturl/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	database.ConnectDB()
	setupRoutes(app)

	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {

}
