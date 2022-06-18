package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	// First way to declare structure, don't use it
	alex := person{"Alex", "Anderson"}
	fmt.Println(alex)

	// Second way to declare structure
	john := person{firstName: "John", lastName: "Smith"}
	fmt.Println(john)

	// Third way to declare structure
	// in this case alice is created with zero values for firstName and lastName
	var alice person

	fmt.Println(alice)
	fmt.Printf("%+v", alice)

	alice.firstName = "Alice"
	alice.lastName = "Wilson"
	fmt.Println(alice)
	fmt.Printf("%+v", alice)
}
