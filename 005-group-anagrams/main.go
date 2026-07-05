// https://leetcode.com/problems/group-anagrams/description/
// Сложность: Medium

package main

import "slices"

func groupAnagrams(strs []string) [][]string {

	mapa := make(map[string][]string)

	for _, v := range strs {

		key := []rune(v)
		slices.Sort(key)

		stringKey := string(key)

		value := mapa[stringKey]
		newArray := append(value, v)
		mapa[stringKey] = newArray

	}

	res := make([][]string, 0, len(mapa))

	for _, v := range mapa {
		res = append(res, v)
	}

	return res
}
