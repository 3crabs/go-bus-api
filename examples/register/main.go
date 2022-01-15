package main

import (
	"context"
	"github.com/3crabs/go-bus-api/bus"
	"github.com/3crabs/go-bus-api/consts"
	"log"
)

func main() {
	b := bus.NewBus("http", consts.Host, "/api/bus")

	err := b.Register(context.Background(), "+7-906-961-25-31")
	if err != nil {
		log.Fatalln(err)
	}
}
