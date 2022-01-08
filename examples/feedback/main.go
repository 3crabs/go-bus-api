package main

import (
	"context"
	"github.com/3crabs/go-bus-api/bus"
	"log"
)

const busHost = "185.119.59.74:8090"

func main() {
	b := bus.NewBus("http", busHost)

	f := bus.FeedbackDTO{
		Phone:   "+7-906-961-25-31",
		Subject: "Проблемы с меню",
		Text:    "Не работает кнопка назад",
	}
	err := b.SendFeedback(context.Background(), f)
	if err != nil {
		log.Fatalln(err)
	}
}
