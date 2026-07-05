// https://leetcode.com/problems/sort-colors/description/
// Сложность: Medium

package main

// алгоритм флага Нидерландов
func sortColors(nums []int) {

	left := 0
	right := len(nums) - 1

	i := 0
	for i <= right {
		switch nums[i] {

		case 0:
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++

		case 1:
			i++

		case 2:
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}
	}
}
