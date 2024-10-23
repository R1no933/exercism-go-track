package elon

import "strconv"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if car.battery <= car.batteryDrain {
		car = car
	} else {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
func (car *Car) DisplayDistance() string {
	return "Driven " + strconv.Itoa(car.distance) + " meters"
}

// TODO: define the 'DisplayBattery() string' method
func (car *Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(car.battery) + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
	if float64(car.battery)/float64(car.batteryDrain) >= float64(trackDistance)/float64(car.speed) {
		return true
	} else {
		return false
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
