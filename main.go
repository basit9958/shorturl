package main

import (
	"fmt"
	"github.com/basit9958/shorturl/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	app := fiber.New()
	router.SetupRoutes(app)

	app.Listen(":3000")
}
