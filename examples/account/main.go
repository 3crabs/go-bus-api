package main

import (
	"context"
	"fmt"
	"github.com/3crabs/go-bus-api/bus"
	"log"
)

const busHost = "185.119.59.74:8090"

func main() {
	b := bus.NewBus("http", busHost)

	login, err := b.Login(context.Background(), "+7-906-961-25-31", "adNJX9")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(login)

	account, err := b.GetAccount(context.Background(), login.AccessToken)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)
}
