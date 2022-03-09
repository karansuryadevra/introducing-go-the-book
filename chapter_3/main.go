// package main

// import "fmt"

// var (
// 	a = 31
// 	b = 43
// )

// func main() {
// 	var x string = "Hello, World!"
// 	fmt.Println(x)

// 	var y = "Short-hand"
// 	fmt.Println(y)

// 	z := "Shortest-hand"
// 	fmt.Println(z)
// }

package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2

	fmt.Println(output)
}
