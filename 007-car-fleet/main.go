// https://leetcode.com/problems/car-fleet/description/
// Сложность: Medium

package main

import "slices"

type Car struct {
	position int
	time     float64
}

func carFleet(target int, position []int, speed []int) int {

	count := 1
	cars := make([]Car, 0, len(position))

	for i := range position {
		cars = append(cars, Car{position: position[i], time: float64(target-position[i]) / float64(speed[i])})
	}

	slices.SortFunc(cars, func(a Car, b Car) int {
		return b.position - a.position
	})

	currTime := cars[0].time

	for i := range len(cars) - 1 {

		if cars[i+1].time > currTime {
			count++
			currTime = cars[i+1].time
		}

	}

	return count
}
