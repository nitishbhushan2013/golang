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

	//Note : range is important as it allow the array to iterate over its element
	// Note : is is the index and str is the current element
	//After each iteration, go throw the index and str and thus we have :=
	// to re-initialize with current index and str.
}
