package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("enter your number:")
	data, _, err := r.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	userNum, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("enter i bit:")
	data, _, err = r.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	iBit, err := strconv.Atoi(string(data))
	if err != nil {
		log.Fatal(err)
	}

	if 1 > iBit || iBit > 64 {
		log.Fatal("i bit must be min 1, max 64")
	}

	iBit--

	fmt.Println("select 0 or 1:")
	data, _, err = r.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	selectedBit, err := strconv.Atoi(string(data))
	if err != nil {
		log.Fatal(err)
	}
	if selectedBit != 0 && selectedBit != 1 {
		log.Fatalf("invalid input: not 0 or 1, your: %v", selectedBit)
	}

	var result int64
	if selectedBit == 0 {
		result = userNum &^ (1 << iBit)
	} else {
		result = userNum | (1 << iBit)
	}

	fmt.Println("____Result____")

	fmt.Printf("your number: %v\nyour i bit: %v\nyour selected bit: %v\nresult: %v\n",
		userNum,
		iBit+1,
		selectedBit,
		result)
}
