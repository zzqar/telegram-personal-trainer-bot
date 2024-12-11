package main

import (
	tele "gopkg.in/telebot.v4"
	"log"
	"telegram-personal-trainer-bot/internal/bot/commands"
	"telegram-personal-trainer-bot/internal/bot/middleware"
	"telegram-personal-trainer-bot/internal/cache/redis"
	"telegram-personal-trainer-bot/internal/client/telegram"
	"telegram-personal-trainer-bot/internal/config"
)

func main() {
	cfg := config.MustConfig()
	rds := redis.MustRedis(cfg.Redis)
	b := telegram.MustBot(cfg.Token)

	c := []tele.Command{
		{Text: "start", Description: "Начать работу с ботом"},
		{Text: "help", Description: "Получить помощь"},
		{Text: "create", Description: "Создать что-то"},
	}
	if err := b.SetCommands(c); err != nil {
		log.Fatalf("Ошибка установки команд: %v", err)
	}

	b.Use(middleware.OnlyMe(*cfg, *rds))

	b.Handle("/start", commands.Start)

	b.Start()
}
