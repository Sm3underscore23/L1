package adapters

import (
	"L1-21/rabbitmq"
)

type rabbitMQAdapter struct {
	rabbitmq.Broker
}

func NewRabbitMQAdapter() rabbitMQAdapter {
	return rabbitMQAdapter{rabbitmq.New()}
}

func (a rabbitMQAdapter) Start() {
	a.StartRabbitMQBroker()
}
func (a rabbitMQAdapter) Proccess() {
	a.RabbitMQProccessing()
}
func (a rabbitMQAdapter) Stop() {
	a.StopRabbitMQBroker()
}
