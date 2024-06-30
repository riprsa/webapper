package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

var _ = godotenv.Load()

var currentWebAppURL string

func main() {
	b, err := tele.NewBot(tele.Settings{
		Token:     os.Getenv("TOKEN"),
		ParseMode: tele.ModeMarkdownV2,
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic(err)
	}

	if os.Getenv("LOG") != "" {
		b.Use(middleware.Logger())
	}

	b.Handle("/start", func(c tele.Context) error {
		if c.Data() != "" {
			currentWebAppURL = c.Data()
		} else {
			if currentWebAppURL == "" {
				return c.Send("usage: you have to specify your Mini App URL with `/start url` first")
			}
		}

		markup := b.NewMarkup()
		markup.Inline(markup.Row(markup.WebApp("open", &tele.WebApp{
			URL: currentWebAppURL,
		})))

		err := c.Send(fmt.Sprintf("current Mini App URL is: `%s`", currentWebAppURL), markup)
		if err != nil {
			return c.Send(fmt.Sprintf("`%s`", err.Error()))
		}
		return nil
	})

	b.Start()
}
