package main

import "fmt"

func variableParams(args ...int) int {
	largest := args[0]
	for _, val := range args {
		if val > largest {
			largest = val
		}
	}
	return largest
}

func main() {
	fmt.Println(variableParams(1, 2, 3, 4, 5, 6, 7, 45, 36, 12))
}
