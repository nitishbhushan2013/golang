package main

import "fmt"

// define custom type
type vehicleList []string

//define function for this type
// any variable of type vehicleList can have access to printVehicle()
// v is reference of current variable with which we are working with
// in this example, v reprersents motorList
func (v vehicleList) printVehicle() {

	for i, vehicle := range v {
		fmt.Println(i, vehicle)
	}
}

// Comparison with Java
// In java, we define class as template which has properties and methods.
// Any objet create from this class has access to these properties and methods
// Class Class { fun1(), fun2()}
// obj1.func1(), obj2.func2()

// go is statically defined language and its not OOPS flavoured
// go has its own way to achieve OOPS
// it defines custom type ehoch actas as a class
// it then define fucntion with receiver to add behaviour to this custom type
// any variable of this custom type has access to this function.
