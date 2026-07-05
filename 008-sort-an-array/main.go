// https://leetcode.com/problems/sort-an-array/
// Сложность: Medium

package main

func quickSort(nums []int, left, right int) {

	if left >= right {
		return
	}

	pivot := nums[(left+right)/2]

	i, j := left, right

	for i <= j {
		for nums[i] < pivot {
			i++
		}
		for nums[j] > pivot {
			j--
		}

		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	quickSort(nums, left, j)
	quickSort(nums, i, right)
}

func sortArray(nums []int) []int {

	quickSort(nums, 0, len(nums)-1)

	return nums
}
