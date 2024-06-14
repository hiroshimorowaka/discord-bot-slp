package main

import (
	"go_bot/bot"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	bot.Run() // call the run function of bot/bot.go
}
