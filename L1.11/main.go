package main

import (
	"fmt"
)

func checker(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	seen := make(map[int]struct{})

	for _, num := range a {
		seen[num] = struct{}{}
	}

	var result []int
	for _, num := range b {
		if _, ok := seen[num]; ok {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3}

	fmt.Println(checker(a, b))
}
