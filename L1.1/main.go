package main

import "fmt"

type wbStudent interface {
	sleep()
	eat()
	code()
}

type human struct{}

type action struct {
	human
}

func (h human) sleep() {
	fmt.Println("[zzzZZZ]")
}

func (h human) eat() {
	fmt.Println("The time has come for this... COOOKIE! AM-NAM-NAM")
}

func (h human) code() {
	fmt.Println("[coding awesome wbshool tasks :3]")
}

func main() {
	p := wbStudent(action{})

	p.sleep()
	p.eat()
	p.code()
}
