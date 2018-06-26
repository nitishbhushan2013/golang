package main

import "fmt"

func addMultiSubstractV1(num1 int, num2 int) (int, int, int) {
	// go invokes type inference [ := ] to craete and assign the value to variable at same line
	add := num1 + num2
	multiply := num1 * num2
	substract := num1 - num2
	return add, multiply, substract
}

// If all the input parameters are of same value then indicate the type at the end
// can define the return value as named parameter in the method signature
func addMultiSubstractV2(num1, num2 int) (add, multiply, substract int) {
	add = num1 + num2
	multiply = num1 * num2
	substract = num1 - num2
	return
}

func closureV1() int {
	i := 0
	return i
}

func closureV2() func() int {
	i := 0
	fmt.Println("i.1 -->", i)
	return func() int {
		i++
		fmt.Println("i.2 -->", i)
		return i

	}
}

func main() {
	fmt.Println(addMultiSubstractV1(5, 3))
	fmt.Println(addMultiSubstractV2(8, 3))
	fmt.Println(closureV1())
	fmt.Println(closureV2())

}
