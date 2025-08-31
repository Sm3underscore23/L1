package main

import (
	"fmt"
	"strings"
)

func reverseSentenceWithBuilder(rowData string) string {
	data := []rune(rowData)
	fmt.Println(string(data))

	b := strings.Builder{}

	j := len(data) - 1
	end := len(data) - 1

	for j > 0 {
		if data[j] == rune(' ') {
			wordLen := end - j
			for i := 1; i <= wordLen; i++ {
				b.WriteRune(data[j+i])
			}
			b.WriteString(" ")
			end = j - 1
		}
		j--
	}

	for i := 0; i <= end; i++ {
		b.WriteRune(data[i])
	}

	return b.String()
}

func reverseSentenceWithSlice(rowData string) string {
	data := strings.Fields(rowData)

	for i, j := 0, len(data)-1; i < j; {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}

	return strings.Join(data, " ")
}

func main() {
	data := "snow dog sun"
	fmt.Println(reverseSentenceWithBuilder(data))
	fmt.Println(reverseSentenceWithSlice(data))
}
