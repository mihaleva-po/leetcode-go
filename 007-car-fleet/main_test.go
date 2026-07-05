package main

import "testing"

func TestCarFleet(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		position []int
		speed    []int
		want     int
	}{
		{
			name:     "classic example 1",
			target:   12,
			position: []int{10, 8, 0, 5, 3},
			speed:    []int{2, 4, 1, 1, 3},
			want:     3,
		},
		{
			name:     "single car",
			target:   10,
			position: []int{3},
			speed:    []int{3},
			want:     1,
		},
		{
			name:     "all merge into one fleet",
			target:   100,
			position: []int{0, 2, 4},
			speed:    []int{4, 2, 1},
			want:     1,
		},
		{
			name:     "no merging, all separate fleets",
			target:   10,
			position: []int{0, 4, 8},
			speed:    []int{1, 1, 1},
			want:     3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := carFleet(tt.target, tt.position, tt.speed)
			if got != tt.want {
				t.Errorf("carFleet(%d, %v, %v) = %d, want %d",
					tt.target, tt.position, tt.speed, got, tt.want)
			}
		})
	}
}