package main

import (
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "classic unsorted",
			input: []int{5, 2, 3, 1},
			want:  []int{1, 2, 3, 5},
		},
		{
			name:  "with duplicates",
			input: []int{5, 1, 1, 2, 0, 0},
			want:  []int{0, 0, 1, 1, 2, 5},
		},
		{
			name:  "already sorted",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "reverse sorted",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "with negative numbers",
			input: []int{-3, 1, -5, 0, 2},
			want:  []int{-5, -3, 0, 1, 2},
		},
		{
			name:  "single element",
			input: []int{7},
			want:  []int{7},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "all same elements",
			input: []int{4, 4, 4, 4},
			want:  []int{4, 4, 4, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortArray(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}