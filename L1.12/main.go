package main

import "fmt"

func set(data []string) []string {
	if len(data) == 0 {
		return nil
	}

	seen := make(map[string]struct{})
	var result []string

	for _, el := range data {
		if _, ok := seen[el]; !ok {
			seen[el] = struct{}{}
			result = append(result, el)
		}
	}

	return result
}

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(set(data))
}
