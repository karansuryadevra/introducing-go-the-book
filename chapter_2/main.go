package main

import "fmt"

func main() {
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("1 + 1 = ", 1.0+1.0)
	fmt.Println("The length of \"Hello, World!\" is ", len("Hello, World!"))
	fmt.Println("The second character of \"Hello, World!\" is ", "Hello, World!"[1])
	fmt.Println("Hello, " + "World!")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
