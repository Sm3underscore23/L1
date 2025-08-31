package main

import (
	"strings"
)

var justString string

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

func someFunc() {
	v := createHugeString(1 << 30)
	justString = string([]byte(v[:100]))

}

func main() {
	someFunc()
}
