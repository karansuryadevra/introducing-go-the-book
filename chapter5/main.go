package main

import "fmt"

func main(){
	var x [5]int
	x[0] = 10
	x[1] = 23
	x[3] = 40
	x[4] = 2

	// Another way to declare an array :-
	// x := [5]float64{
	// 	98,
	// 	32,
	// 	41,
	// }

	total := 0

	for i := 0 ; i <= len(x)-1 ; i ++{
		total += x[i]
	}

	fmt.Println("The test scores are : ", x) // notice how you can print an entire array without using a loop!
	fmt.Println("The average score is ", total / len(x))
	
	total = 0
	for _, value := range x{ // The "_" (underscore) is used to tell the compiler that we are not using a variable here for iterating
		total += value
	}
	fmt.Println("The average score is ", total / len(x))
	
}