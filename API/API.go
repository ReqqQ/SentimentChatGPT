package API

import (
	"cookieChatGPT/API/SentimentService"
	"github.com/gofiber/fiber/v2"
)

type API struct{}
type APIInterface interface {
	GetSentiment(c *fiber.Ctx) string
}

func (api API) GetSentiment(c *fiber.Ctx) string {
	service := SentimentService.Init{}
	return service.GetSentiment(SentimentService.BuildSentimentDTO(c))
}
