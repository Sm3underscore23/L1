package main

import (
	"fmt"
	"strings"
)

func duplicateChecker(data string) bool {
	seen := make(map[string]struct{})
	for _, el := range data {
		el := strings.ToLower(string(el))
		if _, ok := seen[el]; ok {
			return false
		}
		seen[el] = struct{}{}
	}
	return true
}

func main() {
	data := "abcd"

	fmt.Println(duplicateChecker(data))
}
