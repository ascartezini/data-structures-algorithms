package hackerrank

import (
	"testing"
)

func TestMaxSubsetSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int32
		want  int32
	}{
		{
			name:  "test case #1",
			input: []int32{-2, 1, 3, -4, 5},
			want:  8,
		},
		{
			name:  "test case #2",
			input: []int32{2, 1, 5, 8, 4},
			want:  11,
		},
		{
			name:  "test case #3",
			input: []int32{3, 5, -7, 8, 10},
			want:  15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubsetSum(tt.input)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
