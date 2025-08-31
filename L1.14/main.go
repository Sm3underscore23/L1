package main

import (
	"fmt"
	"reflect"
)

func typeChecker(data interface{}) {
	switch data.(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	case bool:
		fmt.Println("bool type")
	default:
		if reflect.TypeOf(data).Kind() == reflect.Chan {
			fmt.Println("chan type")
			return
		}
		fmt.Println("unknown type")
	}
}

func main() {
	var (
		intData        int
		stringData     string
		boolData       bool
		chIntData      chan int
		chStringChData chan string
		chBoolData     chan bool
		chChData       chan chan int
	)

	typeChecker(intData)
	typeChecker(stringData)
	typeChecker(boolData)
	typeChecker(chIntData)
	typeChecker(chStringChData)
	typeChecker(chBoolData)
	typeChecker(chChData)
}
