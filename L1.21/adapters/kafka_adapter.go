package adapters

import "L1-21/kafka"

type kafkAdapter struct {
	kafka.Broker
}

func NewKafkAdapter() kafkAdapter {
	return kafkAdapter{kafka.New()}
}

func (a kafkAdapter) Start() {
	a.StartKafkaBroker()
}
func (a kafkAdapter) Proccess() {
	a.KafkaProccessing()
}
func (a kafkAdapter) Stop() {
	a.StopKafkaBroker()
}
