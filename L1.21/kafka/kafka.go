package kafka

import "fmt"

type Broker interface {
	StartKafkaBroker()
	KafkaProccessing()
	StopKafkaBroker()
}

type kafka struct {
}

func New() kafka {
	return kafka{}
}

func (k kafka) StartKafkaBroker() {
	fmt.Println("*kafka broker started*")
}
func (k kafka) KafkaProccessing() {
	fmt.Println("*kafka processing data*")
}
func (k kafka) StopKafkaBroker() {
	fmt.Println("*kafka broker stopped*")
}
