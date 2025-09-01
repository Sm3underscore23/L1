package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	point "L1.24/point"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("POINT 1\n")
	fmt.Print("x1 = ")
	var x1 float64
	_, err := fmt.Fscan(in, &x1)
	if err != nil {
		log.Fatal(err)
	}
	in.ReadString('\n')

	fmt.Print("y1 = ")
	var y1 float64
	_, err = fmt.Fscan(in, &y1)
	if err != nil {
		log.Fatal(err)
	}
	in.ReadString('\n')

	fmt.Println("-----")

	fmt.Print("POINT 2\n")
	fmt.Print("x2 = ")
	var x2 float64
	_, err = fmt.Fscan(in, &x2)
	if err != nil {
		log.Fatal(err)
	}
	in.ReadString('\n')

	fmt.Print("y2 = ")
	var y2 float64
	_, err = fmt.Fscan(in, &y2)
	if err != nil {
		log.Fatal(err)
	}
	in.ReadString('\n')

	fmt.Println("-----")

	firstPoint := point.NewPoint(x1, y1)
	secondPoint := point.NewPoint(x2, y2)

	fmt.Printf("RESULT\n%v\n", firstPoint.Distance(secondPoint))
}
