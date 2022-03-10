package main

import "fmt"

type Person struct {
	Name string // "Has-a" relationship. A person has a Name
}

func (p *Person) Talk() {
	fmt.Println("Hello, my name is: ", p.Name)
}

type Android struct {
	Person, // "is-a" relationship. An android is a person
	Model string
}

func main() {
	karan := new(Person)
	karan.Talk()

	android18 := new(Android)
	android18.Model = "18"
	android18.Person = "18Name"
	android18.Talk()
}
