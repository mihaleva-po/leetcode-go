package main

import "testing"

func TestSumOfGoodIntegers(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want int
	}{
		{"n=0, k=0 only x=0 fits", 0, 0, 0},
		{"n=1, k=1 mixed matches", 1, 1, 2},
		{"n=2, k=1 single match", 2, 1, 1},
		{"n=5, k=2 no matches", 5, 2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumOfGoodIntegers(tt.n, tt.k)
			if got != tt.want {
				t.Errorf("sumOfGoodIntegers(%d, %d) = %d, want %d", tt.n, tt.k, got, tt.want)
			}
		})
	}
}