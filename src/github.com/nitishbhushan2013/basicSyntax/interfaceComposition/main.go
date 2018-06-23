package main

import (
	"fmt"
)

// Main objectives
// 1. composition over inheritance
// 2. using interface to achieve OOPS concept
// 3. method overloading by interface

// People ...
type People interface {
	Greet()
	ToString()
}

// Address ...
type Address struct {
	addressLine1 string
	suburb       string
	state        string
	pinCode      int
}

// Person ...embedded Address
type Person struct {
	firstName string
	lastName  string
	Address
}

// PrintGreet ...
func PrintGreet(p People) {
	p.Greet()
}

//Greet ...
func (p Person) Greet() {
	fmt.Printf("Hi there!!! I am %s %s", p.firstName, p.lastName)
}

// ToString ...
func (p Person) ToString() {
	fmt.Printf("Hi !!My name is %s %s and I stay at %+v\n", p.firstName, p.lastName, p.Address)
}

// Developer : Since it embeds Person, it is also of type Person and
// since Person has implemented all the interface method, Developer is also
// People
type Developer struct {
	Person
	company  string
	language string
}

//Greet : Overriding Greet() implementation
func (d Developer) Greet() {
	fmt.Printf("Hi, I am developer, working on %s language for %s company",
		d.language, d.company)
}

func main() {
	nitish := Developer{
		Person{"Nitish", "Bhushan", Address{"Artarmon", "Artarmon", "NSW", 2046}}, "sydney Water", "Golang",
	}
	nitish.Greet()
	fmt.Printf("\n\n")
	nitish.ToString()
	fmt.Printf("\n")
}
