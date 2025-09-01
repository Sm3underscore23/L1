package rabbitmq

import "fmt"

type Broker interface {
	StartRabbitMQBroker()
	RabbitMQProccessing()
	StopRabbitMQBroker()
}

type rabbitMQ struct {
}

func New() rabbitMQ {
	return rabbitMQ{}
}

func (r rabbitMQ) StartRabbitMQBroker() {
	fmt.Println("*rabbit_mq broker started*")
}
func (r rabbitMQ) RabbitMQProccessing() {
	fmt.Println("*rabbit_mq processing data*")
}
func (r rabbitMQ) StopRabbitMQBroker() {
	fmt.Println("*rabbit_mq broker stopped*")
}
