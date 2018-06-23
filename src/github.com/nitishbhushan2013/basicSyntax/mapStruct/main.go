package main

import (
	"fmt"
)

// difference between map and struct
/*

struct
	- defined at compile time
	- value could be of any type
	- they are value type - pass by value and hence need to use the pointer
	- can not modify the element dynamically
	- keys do not support indexing
	- use to represent entity with lot of different properties



Map
	- They can define dynamically
	- Keys would be of same type
	- values would be of same type
	- they are reference type, hence no need to implement pointer
	- elements can be defined or updated at run time
	- keys are indexed and we can iterate over index
	- use to represent collection of related properties

*/

func main() {
	// structs defination declaration and use
	type country struct {
		capital  string
		language string
		currency string
	}

	australia := country{
		capital:  "canberra",
		language: "english",
		currency: "aud",
	}
	var india country
	india.capital = "New Delhi"
	india.currency = "rupees"
	india.language = "hindi"

	usa := country{"washington DC", "usd", "english"}

	fmt.Println(australia)
	fmt.Println(india)
	fmt.Println(usa)

	// map definition, declaration and use
	type company map[string]string

	ibm := company{
		"location":   "sydney",
		"department": "AI",
	}
	fmt.Println(ibm)

	oracle := make(map[string]string)
	oracle["department"] = "Data Science"
	oracle["workNature"] = "research"
	oracle["resource"] = "employee"

	fmt.Println(oracle)

	// deleting the key
	delete(oracle, "resource")
	fmt.Println(oracle)

	// iterating the map
	for key, value := range ibm {
		fmt.Println(key, value)
	}

}
