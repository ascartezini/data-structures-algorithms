package hackerrank

import "testing"

// [7, 1, 3, 2, 4, 5, 6]
func TestMinimumSwaps(t *testing.T) {
	tests := []struct {
		name  string
		input []int32
		want  int32
	}{
		{
			name:  "test case #1",
			input: []int32{7, 1, 3, 2, 4, 5, 6},
			want:  5,
		},
		{
			name:  "test case #2",
			input: []int32{2, 3, 4, 1, 5},
			want:  3,
		},
		{
			name:  "test case #3",
			input: []int32{1, 3, 5, 2, 4, 6, 7},
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumSwaps(tt.input)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
