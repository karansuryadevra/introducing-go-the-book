package main

import "fmt"

func main() {
	fmt.Print("Enter the temperature in degree Fahrenheit: ")
	var fahrenheit float32
	fmt.Scanf("%f", &fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9

	fmt.Println("Temperature in Celsius is: ", celsius)
}
