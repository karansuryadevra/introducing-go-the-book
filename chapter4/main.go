package main

import "fmt"

func main(){
	var j int
	for i := 1 ; i <= 10 ; i ++ {
		if i % 2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
		j = i
	}	
		
	switch j {
		case 0: fmt.Println("The number is 0")
		case 10: fmt.Println("The number is 10")
		default: fmt.Println("The number is not 0 or 10")
		}	
}