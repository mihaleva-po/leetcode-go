package main

import (
	"sort"
	"testing"
)

// normalize приводит результат к каноничному виду для сравнения:
// сортирует строки внутри каждой группы и сортирует сами группы.
func normalize(groups [][]string) [][]string {
	normalized := make([][]string, len(groups))
	for i, g := range groups {
		gCopy := make([]string, len(g))
		copy(gCopy, g)
		sort.Strings(gCopy)
		normalized[i] = gCopy
	}
	sort.Slice(normalized, func(i, j int) bool {
		if len(normalized[i]) == 0 {
			return true
		}
		if len(normalized[j]) == 0 {
			return false
		}
		return normalized[i][0] < normalized[j][0]
	})
	return normalized
}

func equalGroups(a, b [][]string) bool {
	na, nb := normalize(a), normalize(b)
	if len(na) != len(nb) {
		return false
	}
	for i := range na {
		if len(na[i]) != len(nb[i]) {
			return false
		}
		for j := range na[i] {
			if na[i][j] != nb[i][j] {
				return false
			}
		}
	}
	return true
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "classic example",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "single empty string",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "single character",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
		{
			name: "no anagrams at all",
			strs: []string{"abc", "def", "ghi"},
			want: [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			if !equalGroups(got, tt.want) {
				t.Errorf("groupAnagrams(%v) = %v, want (up to order) %v", tt.strs, got, tt.want)
			}
		})
	}
}