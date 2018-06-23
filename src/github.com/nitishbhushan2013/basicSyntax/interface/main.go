package main

import (
	"fmt"
)

// Main objectives
// 1. composition over inheritance
// 2. using interface to achieve OOPS concept
// 3. method overloading by interface

//Profession interface : contain method def. uniquely define by each of the implementing type
type Profession interface {
	WorkNature() string
}

//Manager ...
type Manager struct {
}

//Developer ...
type Developer struct {
}

//WorkNature defination for Manager. Its now Profession type .
func (m Manager) WorkNature() string {
	return ("I am manager and my core strength is management")
}

// WorkNature defination for Manager. Its now Profession type.  Method overloading
func (d Developer) WorkNature() string {
	return ("I am developer and my core strength is development")
}

//PrintProfession , any type of Profession can call this function
func PrintProfession(p Profession) {
	fmt.Println(p.WorkNature())
}
func main() {
	manager1 := Manager{}
	developer1 := Developer{}
	PrintProfession(manager1)
	PrintProfession(developer1)

}
