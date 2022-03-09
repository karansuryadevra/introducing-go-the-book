package main

import "fmt"

func evenOdd(number uint) bool {
	halved := number / 2

	fmt.Println(halved)

	if halved%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(evenOdd(99))
}
