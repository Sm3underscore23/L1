package main

import (
	"fmt"
)

func binSearchRec(data []int, el int) bool {
	if 2 > len(data) {
		return data[0] == el
	}
	m := len(data) / 2

	if data[m] > el {
		return binSearchRec(data[:m], el)
	}

	if data[m] < el {
		return binSearchRec(data[m+1:], el)
	}

	return true
}

func binSearchIter(data []int, el int) bool {
	left := 0
	right := len(data) - 1

	for left <= right {
		m := left + (right-left)/2

		if data[m] == el {
			return true
		}

		if data[m] < el {
			left = m + 1
		} else {
			right = m - 1
		}
	}

	return false
}

func main() {
	data := []int{1, 2, 2, 3, 4, 5, 5, 7, 8, 10}

	fmt.Println(binSearchRec(data, 10))
	fmt.Println(binSearchIter(data, 10))
}
