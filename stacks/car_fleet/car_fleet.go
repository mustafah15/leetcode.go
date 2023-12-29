package stacks

import "sort"

func carFleet(target int, position []int, speed []int) int {
	// creating a 2d array to store information about each car
	// car's position and it's speed
	var cars [][2]int = make([][2]int, len(position))

	for i := range cars {
		cars[i] = [2]int{
			position[i],
			speed[i],
		}
	}

	// sorting cars by their positions
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})

	// creating array/stack to store the car fleets
	var fleets []float64 = make([]float64, 0, len(cars))

	// foreach car, we calculate the time required to reach
	// the target position if the calculated time
	// is greater than the time taken
	// by previous car we consider it as fleet
	for i := range cars {
		// calculating current car arrival time
		var arrivalTime float64 = float64(target-cars[i][0]) / float64(cars[i][1])

		// if this is the first car in the fleet
		// or if this car is faster than the car before it
		// larger arrival time means slower car
		if len(fleets) == 0 || arrivalTime > fleets[len(fleets)-1] {
			fleets = append(fleets, arrivalTime)
		}
	}

	return len(fleets)
}
