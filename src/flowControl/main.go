package main

import "fmt"

func initvalue() int {
	return 5
}

func main() {

	// if rules
	// no bracket around if condition
	// curly braces are mandatory
	//curly braces must be on same line of if condition

	if 5 > 0 {
		fmt.Println("number is positive")
	}

	if i := initvalue(); i > 0 {
		fmt.Println("number is positive")
	}

	switch i := initvalue(); {
	case i > 0:
		fmt.Println("Positive number")
	case i < 0:
		fmt.Println("negative number")
	case i == 0:
		fmt.Println("zero number")
	}

}
