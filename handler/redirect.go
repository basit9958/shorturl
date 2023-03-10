package handler

import (
	"github.com/basit9958/shorturl/database"
	"github.com/gofiber/fiber/v2"
)

func RedirectURL(app *fiber.Ctx) error {

	// Short URL from the Params
	url := app.Params("url")

	db1 := database.CreateClient(1)
	defer db1.Close()

	// Get the Long URL from the Short URL
	longURL, err := db1.Get(database.Ctx, url).Result()

	if err != nil {
		return app.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "unable to find the key",
		})
	}
	//	Redirect the request to the long URL
	return app.Redirect(longURL, fiber.StatusPermanentRedirect)

}
