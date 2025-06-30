package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	botToken := os.Getenv("TG_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TG_BOT_TOKEN environment variable is not set")
	}

	// Initialize the Bot
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Enable debug mode for detailed logging
	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Create a keyboard with one button.
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Show my User ID 🆔"),
		),
	)
	keyboard.ResizeKeyboard = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	log.Println("Bot is running and waiting for updates...")

	// Start the loop to process incoming updates.
	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		userID := update.Message.From.ID

		var msg tgbotapi.MessageConfig

		switch update.Message.Text {
		case "/start":
			msg = tgbotapi.NewMessage(chatID, "Hello! 👋 Press the button below to find out your User ID.")
			msg.ReplyMarkup = keyboard
		case "Show my User ID 🆔":
			responseText := fmt.Sprintf("Here is your User ID, %s: `%d`", update.Message.From.FirstName, userID)
			msg = tgbotapi.NewMessage(chatID, responseText)
			msg.ParseMode = tgbotapi.ModeMarkdown
			msg.ReplyMarkup = keyboard
		default:
			msg = tgbotapi.NewMessage(chatID, "Sorry, I don't understand that command. 🤔 Please use the buttons.")
			msg.ReplyMarkup = keyboard
		}

		if _, err := bot.Send(msg); err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}
}
