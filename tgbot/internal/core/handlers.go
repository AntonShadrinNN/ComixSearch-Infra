package core

import (
	"context"
	"fmt"
	"regexp"
	"tgbot/internal/app"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type HandlerGroup struct {
	searcher app.Searcher
}

func (hg HandlerGroup) handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	msg := update.Message.Text
	if !isValidString(msg) {
		b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Keywords must be provided as space-separated values w/o",
				ReplyParameters: &models.ReplyParameters{
					MessageID:                update.Message.ID,
					ChatID:                   update.Message.Chat.ID,
					AllowSendingWithoutReply: true,
				},
			},
		)
		return
	}

	data, err := hg.searcher.GetSuitableComices(ctx, msg, 5)
	if err != nil {
		return
	}

	var resMsg string
	for k, v := range data {
		resMsg += fmt.Sprintf("%s -> %s\n", k, v)
	}

	if resMsg == "" {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "No comices fit your keywords",
		})
		return
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   resMsg,
	})
}

func isValidString(str string) bool {
	pattern := `^[a-zA-Z]+(?:\s+[a-zA-Z]+)*$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(str)
}
