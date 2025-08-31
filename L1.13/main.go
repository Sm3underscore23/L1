package main

import (
	"fmt"
	"math"
)

func simple(a, b int) {
	fmt.Println("simple:")
	fmt.Printf("a: %v\nb: %v\n", a, b)

	if a == b {
		fmt.Printf("a: %v\nb: %v\n", a, b)
		return
	}

	a, b = b, a
	fmt.Printf("a: %v\nb: %v\n", a, b)
}

func xor(a, b int) {

	fmt.Println("XOR:")
	fmt.Printf("a: %v\nb: %v\n", a, b)

	if a == b {
		fmt.Printf("a: %v\nb: %v\n", a, b)
		return
	}

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("a: %v\nb: %v\n", a, b)
}

func arithmetic(a, b int) {

	fmt.Println("arithmetic:")
	fmt.Printf("a: %v\nb: %v\n", a, b)

	if a == b {
		fmt.Printf("a: %v\nb: %v\n", a, b)
		return
	}

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a: %v\nb: %v\n", a, b)
}

func main() {
	a := math.MaxInt64
	b := math.MinInt64

	simple(a, b)
	xor(a, b)
	arithmetic(a, b)
}
