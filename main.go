package main

import (
	"github.com/3crabs/go-bus-api/bus"
)

const busHost = "185.119.59.74:8090"

func main() {
	_ = bus.NewBus("http", busHost)
}
