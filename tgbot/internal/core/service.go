package core

import (
	"context"
	"os"
	"os/signal"
	"tgbot/internal/ports/xkcd"

	"github.com/go-telegram/bot"
)

func Run(botToken string, endpoint string) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	hg := HandlerGroup{
		searcher: xkcd.NewSearcher(ctx, endpoint),
	}
	opts := []bot.Option{
		bot.WithDefaultHandler(hg.handler),
	}

	b, err := bot.New(botToken, opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}
