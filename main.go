package main

import (
	"fmt"
)

var mass float64
var speed float64
var energy float64

func main() {
	fmt.Printf("Enter the objects mass in kilograms: ")
	fmt.Scan(&mass)

	fmt.Printf("Enter the objects speed in meters per second: ")
	fmt.Scan(&speed)

	Calculate(mass, speed)

	fmt.Printf("A object with the mass of %.4f with the velocity of %f will create a total of %f Joules", mass, speed, energy)
}

//Calcuates the amount of energy
func Calculate(mass float64, speed float64) float64 {

	//0.5 * m * v * v
	energy = 0.5 * mass * speed * speed
	return energy
}
