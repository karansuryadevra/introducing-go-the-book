package main

import "fmt"

func main() {
	fmt.Print("Enter the measurement in feet: ")
	var feet float32
	fmt.Scanf("%f", &feet)

	meters := feet / 3.28084
	fmt.Println("The length in meters is: ", meters)
}
