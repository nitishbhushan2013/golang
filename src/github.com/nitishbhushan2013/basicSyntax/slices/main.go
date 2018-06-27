package main

import "fmt"

func main() {

	// Ways to declare slices
	//s1 := []int{1, 2, 3}        // Initialize slice with value
	//s2 := []int{}               // initialize with no value, memoey is allocated
	//var s3 []int                // Declare, memory is not allocated yet
	//s4 := make([]string, 5, 10) // initialize with no value, specify the length and capacity.  memory is allocated

	// length and capacity
	s5 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s5)     // [1 2 3 4 5 6]
	fmt.Println(s5[:4]) // [1 2 3 4]
	fmt.Println(s5[2:]) // [3 4 5 6]

	// s6 still points to slice s5
	s6 := s5[1:4]
	fmt.Println(s6) // [2 3 4]
	fmt.Println("---")
	fmt.Println(len(s6)) //3
	fmt.Println(cap(s6)) //5 --> since s6 still refer to s5, its capacity will be same as s5
	// Note: Above fact points to MEMORY LEAKS, more on memoryLeak.go
	// and hence any changes on s6 would reflect in s5
	s6[2] = 10
	fmt.Println(s6) // [2 3 10]
	fmt.Println(s5) // [1 2 3 10 5 6]

	//append() funcction would take elements or slice and would returns a new slice. This new slice either we can assign to existing sllice or new slice
	

}
