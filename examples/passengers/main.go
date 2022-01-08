package main

import (
	"context"
	"fmt"
	"github.com/3crabs/go-bus-api/bus"
	"log"
	"time"
)

const busHost = "185.119.59.74:8090"

func main() {
	b := bus.NewBus("http", busHost)

	login, err := b.Login(context.Background(), "+7-906-961-25-31", "adNJX9")
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
