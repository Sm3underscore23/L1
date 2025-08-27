package main

import (
	"fmt"
)

func main() {
	data := []int{2, 4, 6, 8, 1}
	fmt.Println("atomic:", withAtomic(data))
	fmt.Println("mutex:", withMutex(data))
	fmt.Println("chan:", withChan(data))
}
