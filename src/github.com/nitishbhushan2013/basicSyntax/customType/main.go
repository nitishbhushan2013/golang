package main

func main() {
	//motorList := []string{}
	motorList := vehicleList{}
	motorList = append(motorList, "Range Rover Sport")
	motorList = append(motorList, "BMW X5")
	motorList = append(motorList, "Audi Q7")

	motorList.printVehicle()

}
