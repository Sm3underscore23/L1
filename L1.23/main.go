package main

import (
	"errors"
	"fmt"
	"log"
)

func deletFromSlice[T any](data []T, idx int) ([]T, error) {
	if idx > len(data)-1 || 0 > idx {
		return nil, errors.New("invalid index")
	}

	result := make([]T, len(data)-1)

	copy(result, data[:idx])
	copy(result[idx:], data[idx+1:])

	return result, nil
}

func main() {
	intSlice := []int{1, 2, 3, 3, 4, 5}
	stringSlice := []string{"cat", "dog", "fish", "humster"}

	intResult, err := deletFromSlice(intSlice, 5)
	if err != nil {
		log.Fatal(err)
	}
	stringResult, err := deletFromSlice(stringSlice, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(intResult)
	intResult[0] = 3
	fmt.Println(stringResult)
}
