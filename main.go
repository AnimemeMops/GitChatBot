package main

import (
	"log"
	"strings"

	"github.com/AnimemeMops/gitChatBot/actions"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	TOKEN     = "1689597486:AAFaVi_kks9rO-jIP3aUXnE4DAklVfDTQcU"
	START_MSG = "Hello! it's Git Bot!"
	SET_TOKEN = "Send your github token. Attention, the message with the token must be in the following format - token:YOUR_TOKEN"
	SET_NAME  = "Send your github name. Attention, the message with the name must be in the following format - name:YOUR_NAME"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}
	var (
		user actions.User
	)
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		var msg tgbotapi.MessageConfig
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch {
		case update.Message.Text == "/start":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, START_MSG)
			bot.Send(msg)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, SET_TOKEN)
			bot.Send(msg)
		case strings.Contains(update.Message.Text, "token:"):
			user.SetToken(strings.Trim(update.Message.Text, "token:"))
			// msg = tgbotapi.NewMessage(update.Message.Chat.ID, actions.Initialization(update.Message.From.UserName, strings.Trim(update.Message.Text, "token:")))
			// bot.Send(msg)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, SET_NAME)
			bot.Send(msg)
		case strings.Contains(update.Message.Text, "name:"):
			user.SetName(strings.Trim(update.Message.Text, "name:"))
			user.Initialization()
		case update.Message.Text == "/repos":
			text, _ := actions.GetAllRepos(&user, strings.Trim(update.Message.Text, "token:"))
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, text)
			bot.Send(msg)
		case strings.Contains(update.Message.Text, "repos:"):
			text := actions.SetRepos(&user, strings.Trim(update.Message.Text, "repos:"))
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, text)
			bot.Send(msg)
		case update.Message.Text == "/commits":
			actions.GetAllCommits(&user)
		}

	}
}
