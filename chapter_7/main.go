package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func circleAreaPointer(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 { // also called a method since it is a function that declares a receiver (c *Circle) as part of its declaration
	return math.Pi * c.r * c.r // use the method to alter the value of the Circle type struct
}

func main() {
	c := Circle{2, 3, 4}
	fmt.Println(circleArea(c))
	fmt.Println(c)

	fmt.Println(circleAreaPointer(&c)) // use pointers with structs so that you can use methods to change the value inside the struct
	fmt.Println(c)                     // the value of c has not changed because we have not modified the value of c in circleAreaPointer()

	fmt.Println(c.area()) // run the area function on c -> is what it's saying
	fmt.Println(c)

}
