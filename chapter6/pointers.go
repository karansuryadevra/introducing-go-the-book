package main

import "fmt"

func zero(xPtr *int){ //xPtr only holds the memory location
	*xPtr = 0	// store the integer 0 in the memory location xPtr refers to
	fmt.Println(xPtr) //will only print the memory location where the calling parameter is stored
}

func one(xPtr *int){
	fmt.Println(*xPtr)
	*xPtr = 1
}

func main(){

	x:=5
	zero(&x) // pass the memory location of the variable x
	fmt.Println(&x) // print the memory location of the variable x
	fmt.Println(x)

	// You can also get a pointer by using "new"

	xPtr := new(int)
	fmt.Println(xPtr)
	one(xPtr)
	fmt.Println(*xPtr)

}