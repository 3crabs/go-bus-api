package main

import (
	"context"
	"fmt"
	"github.com/3crabs/go-bus-api/bus"
	"github.com/3crabs/go-bus-api/consts"
	"log"
)

func main() {
	b := bus.NewBus("https", consts.Host, "")

	login, err := b.Login(context.Background(), consts.Login, consts.Password)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(login)

	paymentTypes, err := b.GetPayments(context.Background(), login.AccessToken)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(paymentTypes)
}
