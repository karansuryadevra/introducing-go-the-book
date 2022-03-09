package main

import "fmt"

func sum(input []int) int {
	total := 0

	for _, val := range input {
		total += val
	}

	return total
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sum(input))
}
