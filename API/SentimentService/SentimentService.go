package SentimentService

import (
	"cookieChatGPT/ChatGPT"
)

type Init struct{}
type SentimentInterface interface {
	GetSentiment(dto *SentimentDTO)
}

func (s Init) GetSentiment(dto *SentimentDTO) string {
	gpt := &ChatGPT.Init{}
	return gpt.Init().GetSentiment(dto.getContent()).Content
}
