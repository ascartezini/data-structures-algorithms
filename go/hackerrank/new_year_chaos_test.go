package hackerrank

import (
	"testing"
)

func TestMinimumBribes(t *testing.T) {
	tests := []struct {
		name  string
		input []int32
		want  string
	}{
		{
			name:  "test case #1",
			input: []int32{1, 2, 5, 3, 7, 8, 6, 4},
			want:  "",
		},
		{
			name:  "test case #2",
			input: []int32{1, 2, 5, 3, 4, 7, 8, 6},
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minimumBribes(tt.input)

			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("got %v, want %v", got, tt.want)
			// }
		})
	}
}
