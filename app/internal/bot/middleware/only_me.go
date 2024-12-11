package middleware

import (
	tele "gopkg.in/telebot.v4"
	"telegram-personal-trainer-bot/internal/cache/redis"
	"telegram-personal-trainer-bot/internal/client/telegram"
	"telegram-personal-trainer-bot/internal/config"
)

func OnlyMe(cfg config.Config, rds redis.StateManager) func(next tele.HandlerFunc) tele.HandlerFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if c.Chat().ID != cfg.AdminId {
				return telegram.SendAndDeletePrevious(c, rds, "Не для публичного пользования")
			}
			return next(c)
		}
	}
}
