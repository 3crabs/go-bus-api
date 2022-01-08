package main

import (
	"github.com/3crabs/go-bus-api/bus"
	"github.com/3crabs/go-bus-api/consts"
)

func main() {
	_ = bus.NewBus("http", consts.Host)
}
