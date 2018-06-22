package main

import "fmt"

// map: it stores statically typed key-value pair. All key have to be of same type
// and all values of the same type.

func main() {
	// first way of creating map
	person := map[string]string{
		"firstName": "Nitish",
		"lastName":  "Bhushan", // weird: we need to put colon even at last element
	}
	fmt.Println(person)

	// second way, using var keyword
	//var car map[string]string // declaration

	// Third way of declaring map
	car := make(map[string]string)

	// populating map, must use []
	car["make"] = "LandRover"
	car["model1"] = "RangeRover Evoque"
	car["model2"] = "RangeRover Velar"
	car["model3"] = "RangeRover Sport"
	car["model4"] = "Discovery"
	fmt.Println(car)

	//delete
	delete(car, "model4")
	fmt.Println(car)

	//iterating
	fmt.Println("******* Iterating the map *******")
	for key, value := range car {
		fmt.Println(key, value)
	}

}
