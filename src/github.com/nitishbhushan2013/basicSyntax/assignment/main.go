package main // two conditions
// 1. its executable file
// 2. func name must be main()
import "fmt" // helper module

func main() {
	// assignment
	// 1. long form
	var name1 = "long way of creating variables"
	//2. short form, used only at initialization time
	name2 := "intiutive way: let compiler figure out based on ':=' and right hand assignemnt"
	number1 := 123
	flag := true
	//Note: name declared must be used in the program otherwise compiler error
	//Note : if variable is not being used then replace it with _

	fmt.Println("output ==>", name1, name2, number1, flag)
}
