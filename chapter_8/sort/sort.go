package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (byName ByName) Len() int {
	return len(byName)
}

func (byName ByName) Less(i, j int) bool {
	return byName[i].Name < byName[j].Name
}

func (byName ByName) Swap(i, j int) {
	byName[i], byName[j] = byName[j], byName[i]
}

func main() {
	kids := []Person{
		{"Shruti", 25},
		{"Karan", 25},
	}

	sort.Sort(ByName(kids))

	fmt.Println(kids)
}
