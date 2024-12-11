package telegram

import (
	"fmt"
	tele "gopkg.in/telebot.v4"
	"log"
	"strconv"
	"telegram-personal-trainer-bot/internal/cache/redis"
	"time"
)

func MustBot(token string) *tele.Bot {
	b, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalf("Error connect to telegram: %v", err)
	}
	return b
}

func SendAndDeletePrevious(c tele.Context, rds redis.StateManager, text string) error {
	userKey := fmt.Sprintf("last_msg_id:%d", c.Sender().ID)
	// Удаляем предыдущее сообщение
	lastMsgID, err := rds.Get(userKey)
	if err == nil {
		_ = c.Bot().Delete(tele.StoredMessage{
			MessageID: lastMsgID,
			ChatID:    c.Chat().ID,
		})
	}
	// Отправляем новое сообщение
	msg, err := c.Bot().Send(c.Sender(), text)
	if err != nil {
		return err
	}
	return rds.Set(userKey, strconv.Itoa(msg.ID))
}
