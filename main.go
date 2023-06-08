package main

import (
	"cookieChatGPT/Routes"
	"cookieChatGPT/Server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	app := Server.Init()
	Routes.InitBase(app)
	Routes.InitPostRoutes(app)
	Server.Start(app)
}
