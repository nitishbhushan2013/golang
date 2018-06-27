package main

import "fmt"

// This example demonstarte the memory leak in slice and way to overcome it
func memoryLeak() []int {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := s1[1:5]
	fmt.Println(s2)
	return s2
}

// copy(target, source) would make a different array for the slice and hence both slice are independen
func memoryLeakSolution() []int {
	s11 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s22 := make([]int, 4)
	copy(s22, s11[1:5])
	return s22
}

func main() {
	s3 := memoryLeak()
	fmt.Println("In the main code ")
	fmt.Println("s3 array->", s3)                   //[2 3 4 5]
	fmt.Println(len(s3))                            //4
	fmt.Println(cap(s3))                            //8
	fmt.Println("underlying array->", s3[:cap(s3)]) //[2 3 4 5 6 7 8 9]

	fmt.Println("-----------------------------------------------------------------")
	s33 := memoryLeakSolution()
	fmt.Println("s33 array->", s33)                   //[2 3 4 5]
	fmt.Println(len(s33))                             //4
	fmt.Println(cap(s33))                             //4
	fmt.Println("underlying array->", s33[:cap(s33)]) //[2 3 4 5]

}
