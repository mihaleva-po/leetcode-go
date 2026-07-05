// https://leetcode.com/problems/valid-binary-strings-with-cost-limit/
// Сложность: Medium

package main

import (
	"math"
	"strconv"
	"strings"
)

func generateValidStrings(n int, k int) []string {

	maxNum := int(math.Pow(2, float64(n)))

	res := []string{}

	for v := range maxNum {

		b := ""

		for v > 0 {
			ostatok := v % 2
			b = strconv.Itoa(ostatok) + b
			v /= 2
		}

		binary := strings.Repeat("0", n-len(b)) + b

		cost := 0
		var last string
		for i := range binary {

			v := binary[i : i+1]

			if v == "1" && v == last {
				cost = 10000
			}

			if v == "1" {
				cost += i
			}

			last = v
		}

		if cost <= k {
			res = append(res, binary)
		}

	}

	return res
}
