package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"receiptbot/handlers"
)

func NewBotAPI(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot.Debug = true
	return bot, nil
}
func Start(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Photo != nil || update.Message.Document != nil {
			handlers.HandlePhoto(bot, update.Message)
		} else if update.Message.Text != "" {
			handlers.HandleText(bot, update.Message)
		}
	}
}
