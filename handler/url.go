package handler

import (
	"github.com/basit9958/shorturl/database"
	"github.com/basit9958/shorturl/models"
	"github.com/gofiber/fiber/v2"
)

func GetURL(app *fiber.Ctx) error {

	// Short URL from the Params
	url := app.Params("url")

	db1 := database.CreateClient(1)
	defer db1.Close()

	// Get the Long URL from the Short URL
	longURL, err := db1.Get(database.Ctx, url).Result()

	if err != nil {
		return err
	}

	reponsebody := models.Response{
		URL:      longURL,
		ShortUrl: url,
	}
	return app.JSON(reponsebody)
}
