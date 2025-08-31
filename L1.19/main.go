package main

import (
	"fmt"
)

func reverseWord(data string) string {
	if 2 > len(data) {
		return data
	}

	result := []rune(data)

	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}

	return string(result)
}

func main() {
	data := "главрыба"
	fmt.Println(reverseWord(data))
}
