package ChatGPT

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type RequestBody struct {
	Model       string     `json:"model"`
	Messages    []Messages `json:"messages"`
	Temperature float32    `json:"temperature"`
	N           int        `json:"n"`
}

const CHATGTP_COMPLETIONS_ENDPOINT = "https://api.openai.com/v1/chat/completions"

func Init() *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	req.Header.Set("Authorization", "Bearer ")
	req.Header.Set("Content-Type", "application/json")

	return req
}
func GetSentiment(req *fasthttp.Request, message string) {
	req.Header.SetMethod("POST")
	req.SetRequestURI(CHATGTP_COMPLETIONS_ENDPOINT)

	requestBody := RequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []Messages{
			{Role: "user", Content: "What is sentiment(positive,negative,neutral) of this message: '" + message + "'"},
		},
		Temperature: 1,
		N:           1,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {

	}
	req.SetBody(jsonData)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := fasthttp.Do(req, resp); err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Printf("Response body: %s\n", resp.Body())
}
