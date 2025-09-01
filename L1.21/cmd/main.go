package main

import (
	"L1-21/adapters"
)

type brocker interface {
	Start()
	Proccess()
	Stop()
}

func main() {
	k := brocker(adapters.NewKafkAdapter())

	k.Start()
	k.Proccess()
	k.Stop()

	r := brocker(adapters.NewRabbitMQAdapter())

	r.Start()
	r.Proccess()
	r.Stop()
}
