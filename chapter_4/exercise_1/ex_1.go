package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			if i%6 == 0 {
				fmt.Println(i)
			}
		}
	}
}
