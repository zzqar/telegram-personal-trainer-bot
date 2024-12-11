package commands

import (
	tele "gopkg.in/telebot.v4"
)

var (
	menu      = &tele.ReplyMarkup{}
	btnChoose = menu.Text("Выбрать программу")
	btnCreate = menu.Text("Создать программу")
)

func Start(c tele.Context) error {
	menu.ResizeKeyboard = true
	menu.Reply(
		menu.Row(btnChoose),
		menu.Row(btnCreate),
	)
	return c.Send("Выберите действие:", menu)
}
