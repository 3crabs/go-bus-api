package main

import (
	"context"
	"fmt"
	"github.com/3crabs/go-bus-api/bus"
	"github.com/3crabs/go-bus-api/consts"
	"log"
)

func main() {
	b := bus.NewBus("http", consts.Host, "/api/bus")

	login, err := b.Login(context.Background(), consts.Login, consts.Password)
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
