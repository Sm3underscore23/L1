package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("a = ")
	data, _, err := r.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	a := new(big.Rat)
	_, ok := a.SetString(string(data))
	if !ok {
		log.Fatal("a invalid number")
	}

	fmt.Print("b = ")
	data, _, err = r.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	b := new(big.Rat)
	_, ok = b.SetString(string(data))
	if !ok {
		log.Fatal("b invalid number")
	}

	fmt.Print("choose operation (\"+\", \"-\", \"*\", \"/\"): ")
	operation, _, err := r.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	result := new(big.Rat)

	switch string(operation) {
	case "+":
		result.Add(a, b)
	case "-":
		result.Sub(a, b)
	case "*":
		result.Mul(a, b)
	case "/":
		if b.Sign() == 0 {
			log.Fatal("division by zero")
		}
		result.Quo(a, b)
	default:
		log.Fatal("invalid operation")
	}

	s := strings.TrimRight(result.FloatString(10), "0")
	s = strings.TrimRight(s, ".")

	fmt.Println("result:", s)
}
