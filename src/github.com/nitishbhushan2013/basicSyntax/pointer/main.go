package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	person1 := person{
		firstName: "Nitish",
		lastName:  "Bhushan"}

	fmt.Println(person1.firstName)
	person1.updateFirstName("Manish") // this will not change the firstName

	// use pointer to achieve the same
	// first approach
	personPointer := &person1
	personPointer.updateName("Jolly")

	// second approach
	person1.updateName("Molly")

	fmt.Println(person1.firstName)

}

// will not change the firstName
func (p person) updateFirstName(newName string) {
	p.firstName = newName
}

func (pointer *person) updateName(newName string) {
	(*pointer).firstName = newName
}
