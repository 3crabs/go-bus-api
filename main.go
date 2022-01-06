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

	from, err := b.GetPointsFrom(context.Background(), "барнаул")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("from", from)

	to, err := b.GetPointsTo(context.Background(), (*from)[0].Id, "новосибирск")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("to", to)

	races, err := b.GetRaces(context.Background(), (*from)[0].Id, (*to)[0].Id, "08.01.2022", 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(races)
}
