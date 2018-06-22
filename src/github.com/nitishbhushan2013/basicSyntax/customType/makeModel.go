package main

import "fmt"

type model []string
type make []string

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
func populateModel(s string) model {
	modelList := model{}
	switch s {
	case "BMW":
		modelList = append(modelList, "X1", "X2", "X3", "X4", "X5", "X6")
	case "LandRover":
		modelList = append(modelList, "Evoque", "Velar", "Sport")

	}
	return modelList
}

func main() {

	fmt.Println("Creating list of car  make")
	createMakeList().print()
	fmt.Println("Creating model list for 'BMW'")
	bmwModelList := populateModel("BMW")
	fmt.Println("Printing list for 'BMW' models")
	bmwModelList.print()

	fmt.Println("Creating model list for 'landRover'")
	landRoverList := populateModel("LandRover")

	fmt.Println("Printing list for 'landRover' models")
	landRoverList.print()

}
