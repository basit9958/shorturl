package router

import (
	"github.com/basit9958/shorturl/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/get/:url", handler.GetURL)
	app.Post("/api", handler.URLShortner)

}
