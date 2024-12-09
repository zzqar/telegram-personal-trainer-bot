package main

import (
	tele "gopkg.in/telebot.v4"
	"log"
	"telegram-personal-trainer-bot/internal/cache/redis"
	"telegram-personal-trainer-bot/internal/config"
	"time"
)

func main() {
	cfg := config.MustConfig()
	_ = redis.MustRedis(cfg.Redis)

	pref := tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
