package Routes

import (
	"cookieChatGPT/API"
	"github.com/gofiber/fiber/v2"
)

func InitBase(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})
}
func InitPostRoutes(app *fiber.App) {
	app.Post("/sentiment", func(c *fiber.Ctx) error {
		api := API.API{}

		return c.JSON(api.GetSentiment(c))
	})
}
