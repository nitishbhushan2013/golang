package main

import "fmt"

type model []string
type make []string
type carMake string

// receiver function for model type
func (m model) print() {
	for i, car := range m {
		fmt.Println(i, car)
	}
}

// receiver function for make type
func (m make) print() {
	for i, car := range m {
		fmt.Println(i, car)
	}
}

// return of make type
func createMakeList() make {
	makeList := make{"BMW", "landRover", "Hummer"}
	return makeList
}

//return of model type
func (car carMake) populateModel() model {
	modelList := model{}
	switch car {
	case "BMW":
		modelList = append(modelList, "X1", "X2", "X3", "X4", "X5", "X6")
	case "landRover":
		modelList = append(modelList, "Evoque", "Velar", "Sport")

	}
	return modelList
}

func main() {

	fmt.Println("Available car make are --> ")
	createMakeList().print()

	for _, car := range createMakeList() { // replacing i with _ as it is not being used
		c := carMake(car)
		fmt.Println("printing list of model for make", string(c))
		fmt.Println(c.populateModel())
	}
}
