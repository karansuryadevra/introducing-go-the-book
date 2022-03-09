package main

import "fmt"

func main() {

	i := 1

	fmt.Println("Print1: ")

	for i <= 10 {
		fmt.Print(i, " ")
		i = i + 1

		switch i {
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		case 6:
			fmt.Println("Six")
		case 7:
			fmt.Println("Seven")
		case 8:
			fmt.Println("Eight")
		case 9:
			fmt.Println("Nine")
		case 10:
			fmt.Println("Ten")
		case 1:
			fmt.Println("One")

		}
	}

	fmt.Println(" \nPrint2: ")

	for i := 1; i <= 10; i++ {
		fmt.Print(i, "-")
		if i%2 == 0 {
			fmt.Print("Even ")
		} else {
			fmt.Print("Odd ")
		}
	}

}
