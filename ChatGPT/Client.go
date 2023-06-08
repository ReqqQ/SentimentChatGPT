package ChatGPT

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

const BASE_MODEL = "gpt-3.5-turbo"
const REQUEST_TYPE_POST = "POST"

type Init struct {
	Client *fasthttp.Request
}

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

func (gpt *Init) Init() Init {
	gpt.Client = fasthttp.AcquireRequest()
	gpt.Client.Header.Set("Authorization", "Bearer "+os.Getenv("CHAT_GPT_KEY"))
	gpt.Client.Header.Set("Content-Type", "application/json")

	return *gpt
}

func (gpt *Init) setRequestParams(requestBody RequestBody) []byte {
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
	}
	gpt.Client.SetBody(jsonData)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := fasthttp.Do(gpt.Client, resp); err != nil {
		log.Fatalf("Error: %s", err)
	}

	return resp.Body()
}
