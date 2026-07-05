package main

import "testing"

func TestValidPath(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		edges       [][]int
		source      int
		destination int
		want        bool
	}{
		{
			name:        "path exists via triangle",
			n:           3,
			edges:       [][]int{{0, 1}, {1, 2}, {2, 0}},
			source:      0,
			destination: 2,
			want:        true,
		},
		{
			name:        "disconnected components",
			n:           6,
			edges:       [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}},
			source:      0,
			destination: 5,
			want:        false,
		},
		{
			name:        "source equals destination, no edges needed",
			n:           1,
			edges:       [][]int{},
			source:      0,
			destination: 0,
			want:        true,
		},
		{
			name:        "isolated node with no edges",
			n:           2,
			edges:       [][]int{},
			source:      0,
			destination: 1,
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validPath(tt.n, tt.edges, tt.source, tt.destination)
			if got != tt.want {
				t.Errorf("validPath(%d, %v, %d, %d) = %v, want %v",
					tt.n, tt.edges, tt.source, tt.destination, got, tt.want)
			}
		})
	}
}