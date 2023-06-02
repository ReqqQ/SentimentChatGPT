package main

import (
	"bufio"
	"cookieChatGPT/ChatGPT"
	"os"
)

func main() {
	//sentimentList := [3]string{"positive", "negative", "neutral"}
	println("Sentiment analyse engine started ! ")
	println("Enter your message: ")
	message := getMessageInput()
	client := ChatGPT.Init()
	ChatGPT.GetSentiment(client, message)
}
func getMessageInput() string {
	message := bufio.NewScanner(os.Stdin)
	message.Scan()

	return message.Text()
}
