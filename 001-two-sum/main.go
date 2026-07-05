// https://leetcode.com/problems/two-sum/
// Сложность: Easy

package main

func twoSum(nums []int, target int) []int {

mapa := make(map[int]int, len(nums))

for i, v := range nums {
    if index, ok := mapa[target - v]; ok {
return []int{i, index}
    }
    mapa[v] = i
}

 return []int{}
}