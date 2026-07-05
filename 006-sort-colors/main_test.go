package main

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "classic example",
			input: []int{2, 0, 2, 1, 1, 0},
			want:  []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:  "already sorted",
			input: []int{0, 1, 2},
			want:  []int{0, 1, 2},
		},
		{
			name:  "reverse sorted",
			input: []int{2, 1, 0},
			want:  []int{0, 1, 2},
		},
		{
			name:  "single element",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "all same color",
			input: []int{1, 1, 1, 1},
			want:  []int{1, 1, 1, 1},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.input)
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("sortColors() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}