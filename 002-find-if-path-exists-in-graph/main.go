// https://leetcode.com/problems/find-if-path-exists-in-graph
// Сложность: Easy

package main

func checkNode(node, destination int, matrix map[int][]int, used map[int]struct{}) bool {
	if node == destination {
		return true
	}
	value, ok := matrix[node]

	if !ok {
		return false
	}
	for _, v := range value {
		if _, isInclide := used[v]; !isInclide {
			used[v] = struct{}{}
			res := checkNode(v, destination, matrix, used)
			if res {
				return true
			}
		}

	}

	return false
}

func validPath(n int, edges [][]int, source int, destination int) bool {

	matrix := make(map[int][]int, n)

	for _, edge := range edges {

		value, ok := matrix[edge[0]]
		if ok {
			updateValue := append(value, edge[1])
			matrix[edge[0]] = updateValue
		} else {
			matrix[edge[0]] = []int{edge[1]}
		}

		value, ok = matrix[edge[1]]
		if ok {
			updateValue := append(value, edge[0])
			matrix[edge[1]] = updateValue
		} else {
			matrix[edge[1]] = []int{edge[0]}
		}
	}

	return checkNode(source, destination, matrix, map[int]struct{}{})

}
