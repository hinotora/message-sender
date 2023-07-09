package main

import (
	"flag"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Command struct {
	Token 		string
	Chat 		int64
	Text   		string
}

func ParseCLI() *Command {
	command := Command{}

	flag.StringVar(&command.Token, "token", "", "Bot token from TelegramAPI (check at @BotFather)")

	flag.Int64Var(&command.Chat, "chat", 0, "Chat identifier (user id)")

	flag.StringVar(&command.Text, "text", "", "Text of message")

	flag.Parse()

	return &command
}

func main() {
	command := ParseCLI()

	bot, _ := tgbotapi.NewBotAPI(command.Token)

	msg := tgbotapi.NewMessage(command.Chat, command.Text)

	_, err := bot.Send(msg)

	if err != nil {
		fmt.Printf("FAIL. %s", err.Error())
	} else {
		fmt.Println("OK. Message send.")
	}
}
