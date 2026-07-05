// https://leetcode.com/problems/sum-of-compatible-numbers-in-range-i
// Сложность: Easy

package main

import "math"

func sumOfGoodIntegers(n int, k int) int {

	sum := 0

	maxX := k + n

	for x := 0; x <= maxX; x++ {
		if (int(math.Abs(float64(n-x))) <= k) && (n&x == 0) {
			sum += x
		}
	}

	return sum

}