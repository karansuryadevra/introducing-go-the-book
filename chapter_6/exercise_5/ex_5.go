package main

import "fmt"

func fibo(num int) int {
	if num == 1 {
		return 1
	}
	return num + fibo(num-1)
}

func main() {
	fmt.Println(fibo(5))
}
