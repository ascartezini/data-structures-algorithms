package hackerrank

import (
	"testing"
)

func TestJumpingOnClouds(t *testing.T) {
	tests := []struct {
		name  string
		input []int32
		want  int32
	}{
		{
			name:  "test case #1",
			input: []int32{0, 1, 0, 0, 0, 1, 0},
			want:  3,
		},
		{
			name:  "test case #2",
			input: []int32{0, 0, 0, 1, 0, 0},
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jumpingOnClouds(tt.input)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
