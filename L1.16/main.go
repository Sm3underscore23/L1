package main

import (
	"fmt"
)

func quickSort(data []int) []int {
	if 2 > len(data) {
		return data
	}

	pivot := len(data) / 2

	left, right := 0, len(data)-1

	data[pivot], data[right] = data[right], data[pivot]

	for i := range data {
		if data[i] < data[right] {
			data[i], data[left] = data[left], data[i]
			left++
		}
	}

	data[right], data[left] = data[left], data[right]

	quickSort(data[:left])
	quickSort(data[left+1:])

	return data
}

func main() {
	s := []int{8, 2, 5, 2, 5, 3, 7, 10, 4, 1}

	fmt.Println(quickSort(s))
}
