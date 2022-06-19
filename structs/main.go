package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@example.com",
			zipCode: 94000,
		},
	}

	jim.print()

	// & - give us the memory address of the value this variable is pointing at
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	// New name is going to be displayed
	jim.print()

	// New example that works differently with structs
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	// Here we could see that the first element of slice was changed
	// It happens because slice is a reference type
	// Under the hood when we create a new slice, we create array and a special object
	// this special object has a pointer to the original array and other fields
	// In reality updateSlice passes special object that refers to orignal array.
	// That is why here we could see updated array
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

// *person - this is a type description - it means we're worikng
// with a pointer to a person
// * - is not operator here, it's a description of type
func (pointerToPerson *person) updateName(newName string) {
	// * - give us the value this memory address is pointing at
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
