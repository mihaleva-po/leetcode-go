package main

import (
	"reflect"
	"testing"
)

func TestGenerateValidStrings(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want []string
	}{
		{
			name: "n=2, k=0 excludes any 1 not at index 0",
			n:    2,
			k:    0,
			want: []string{"00", "10"},
		},
		{
			name: "n=2, k=1 allows single 1 at index 0 or 1",
			n:    2,
			k:    1,
			want: []string{"00", "01", "10"},
		},
		{
			name: "n=1, k=0 both single-bit strings fit",
			n:    1,
			k:    0,
			want: []string{"0", "1"},
		},
		{
			name: "n=3, k=1 filters out consecutive ones and high-cost strings",
			n:    3,
			k:    1,
			want: []string{"000", "010", "100"},
		},
		{
			name: "n=0, k=0 only the empty string",
			n:    0,
			k:    0,
			want: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateValidStrings(tt.n, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateValidStrings(%d, %d) = %v, want %v", tt.n, tt.k, got, tt.want)
			}
		})
	}
}
