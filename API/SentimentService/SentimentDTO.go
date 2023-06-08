package SentimentService

import "github.com/gofiber/fiber/v2"

type SentimentDTO struct {
	Content string `json:"sentiment" form:"sentiment"`
}

func BuildSentimentDTO(c *fiber.Ctx) *SentimentDTO {
	dto := new(SentimentDTO)

	err := c.BodyParser(dto)
	if err != nil {
		return nil
	}

	return dto
}

func (dto SentimentDTO) getContent() string {
	return dto.Content
}
