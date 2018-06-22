package main

import "fmt"

func main() {

	// array stores only one type
	sliceArr := []string{}                                       // dynamic array : it can grow or shrink
	staticArr := []string{"one", "two", "three", "four", "five"} // its static

	// adding element in dynamic array
	sliceArr = append(sliceArr, "step1") // it does not update current slice but returns new slice

	fmt.Println(staticArr)
	fmt.Println(sliceArr)

	// iterating the element
	for i, str := range staticArr {
		fmt.Println(i, str)
	}
}
