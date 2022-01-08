package main

import (
	"context"
	"github.com/3crabs/go-bus-api/bus"
	"github.com/3crabs/go-bus-api/consts"
	"log"
)

func main() {
	b := bus.NewBus("http", consts.Host)

	f := bus.FeedbackDTO{
		Phone:   "+7-000-000-00-00",
		Subject: "Проблемы с меню",
		Text:    "Не работает кнопка назад",
	}
	err := b.SendFeedback(context.Background(), f)
	if err != nil {
		log.Fatalln(err)
	}
}
