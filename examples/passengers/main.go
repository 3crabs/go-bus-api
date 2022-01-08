package main

import (
	"context"
	"fmt"
	"github.com/3crabs/go-bus-api/bus"
	"github.com/3crabs/go-bus-api/consts"
	"log"
	"time"
)

func main() {
	b := bus.NewBus("http", consts.Host)

	login, err := b.Login(context.Background(), consts.Login, consts.Password)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(login)

	passengers, err := b.GetPassengers(context.Background(), login.AccessToken)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(passengers)

	for _, p := range *passengers {
		if !p.Owner {
			err = b.DeletePassenger(context.Background(), login.AccessToken, p.Id)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	p := bus.PassengerCreateDTO{
		Birthday:    time.Time{},
		Citizenship: "",
		DocNum:      "",
		DocSeries:   "",
		DocTypeCode: "",
		Email:       "",
		FirstName:   "",
		Gender:      "",
		LastName:    "",
		MiddleName:  "",
		Owner:       false,
		Phone:       "",
	}
	passenger, err := b.AddPassenger(context.Background(), login.AccessToken, p)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(passenger)

	err = b.DeletePassenger(context.Background(), login.AccessToken, passenger.Id)
	if err != nil {
		log.Fatalln(err)
	}

	passengers, err = b.GetPassengers(context.Background(), login.AccessToken)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(passengers)
}
