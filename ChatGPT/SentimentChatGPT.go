package ChatGPT

import (
	"encoding/json"
	"log"
)

const CHATGPT_COMPLETIONS_ENDPOINT = "https://api.openai.com/v1/chat/completions"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type SentimentResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func (s SentimentResponse) getFirstChoice() Message {
	return s.Choices[0].Message
}

func getRequestBody(message string) RequestBody {
	return RequestBody{
		Model: BASE_MODEL,
		Messages: []Messages{
			{Role: "user", Content: "Sentiment of this message:" + message + ".Response only in one word (positive,negative,neutral)"},
		},
		Temperature: 1,
		N:           1,
	}
}

func decodeResponse(response []byte) Message {
	var sentimentResponse SentimentResponse
	err := json.Unmarshal(response, &sentimentResponse)
	if err != nil {
		log.Fatal(err)
	}

	return sentimentResponse.getFirstChoice()
}

func (gpt Init) GetSentiment(message string) Message {
	gpt.Client.Header.SetMethod(REQUEST_TYPE_POST)
	gpt.Client.SetRequestURI(CHATGPT_COMPLETIONS_ENDPOINT)

	response := gpt.setRequestParams(getRequestBody(message))

	return decodeResponse(response)
}
