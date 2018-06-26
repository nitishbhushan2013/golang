package main

import "fmt"

// go supports 'pass by value'. There are two type - reference type and value type
// All the primitives are value type.
//we can use pointer to change the value of primitive type which is pass to function

func changePrimitiveValue(num1 *int) int {
	*num1 = 10
	return *num1
}

func main() {
	num1 := 3        // define num1, which is ideally stored at some address in RAM
	var p1 *int      // define pointer to type int
	p1 = &num1       // assign memory address of num1 to p1, it will point to the num1 memory address
	fmt.Println(p1)  // print the memory address
	fmt.Println(*p1) // dereference the pointer to get the value stored at this memory address. get the value,					 this pointer points to
	fmt.Println(&p1) // get the memory address of the location which store 'memory address of num1'

	// type inference
	num2 := 6
	p2 := &num2
	fmt.Println(*p2)

	// //we can use pointer to change the value of primitive type which is pass to function
	fmt.Println("initial value of function parameter is 5")
	i := 5
	p3 := &i
	fmt.Println("final value is ", changePrimitiveValue(p3))

}
