package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("512657781:AAH3BTfWk0gas92YmfNnCv_6n8RdJUY2cOM")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		msg.Text = update.Message.Text

		db, err := gorm.Open("postgres",
			"host=localhost user=rsuh1 dbname=opencart sslmode=disable password=qwer1234")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		p := Post{Body: msg.Text}

		db.Debug().Save(&p)

		bot.Send(msg)

	}
}

type Post struct {
	//Title  string
	Body string
}
