package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(
	acceletarion, initial_velocity, initial_displacement float64,
) func(float64) float64 {
	displacement_function := func(time float64) float64 {
		var result float64
		result = (acceletarion * (math.Pow(time, 2)) / float64(
			2,
		)) + (initial_velocity * time) + initial_displacement
		return result
	}
	return displacement_function
}

func main() {
	var acceleration float64
	var initial_velocity float64
	var initial_displacement float64
	var time float64

	fmt.Println("Acceleration:")
	fmt.Scan(&acceleration)
	fmt.Println("Initial Velocity:")
	fmt.Scan(&initial_velocity)
	fmt.Println("Initial Displacement:")
	fmt.Scan(&initial_displacement)
	fmt.Println("Time:")
	fmt.Scan(&time)

	displacement_function := GenDisplaceFn(acceleration, initial_velocity, initial_displacement)
	fmt.Println("Displacement:")
	fmt.Println(displacement_function(time))
}
