package main

import (
	"fmt"
)

type departmentDetail struct {
	deptName string
	empCount int
}

type employee struct {
	firstName string
	lastName  string
	company   string
	//deptDetail departmentDetail  // one way of defining embedded struct
	//departmentDetail
}

func main() {
	// first way of initializing struct
	emp1 := employee{"nitish", "bhushan", "sydney water"}
	//second way of initializing struct, json format
	emp2 := employee{
		firstName: "nitish",
		lastName:  "bhushan",
		company:   "sydney water",
		/*departmentDetail : departmentDetail {
			deptName : "marketing",
			empCount : 500
		}*/
	}
	//third way of initializing struct, json format
	var emp3 employee
	emp3.firstName = "nitish"
	emp3.lastName = "bhushan"
	emp3.company = "sydney water"

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
}
