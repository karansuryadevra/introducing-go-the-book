package main

import "fmt"

func swap(x, y *int) {
	temp := 0
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println("x is: ", x, " ", "y is: ", y)
}
