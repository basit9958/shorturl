package handler

import (
	"github.com/basit9958/shorturl/database"
	"github.com/basit9958/shorturl/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func URLShortner(app *fiber.Ctx) error {

	requestBody := new(models.Request)

	if err := app.BodyParser(requestBody); err != nil {
		return app.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot Parse the Request ",
		})
	}

	db1 := database.CreateClient(1)
	defer db1.Close()

	// Generate a random ID of 6 char if URL_short is not there
	var rndId string
	if requestBody.ShortUrl == "" {
		rndId = uuid.New().String()[:6]
	} else {
		rndId = requestBody.ShortUrl
	}

	// Check Whether the Short URL is already present in the database or not
	// Possible value Empty or String
	longURL, err := db1.Get(database.Ctx, rndId).Result()
	if err != nil {
		return err
	}

	if longURL != "" {
		return app.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": " URL is already in the use",
		})
	}

	return nil

}

func createShort() {

}
