package hackerrank

import (
	"reflect"
	"testing"
)

func TestRotationLeft(t *testing.T) {
	tests := []struct {
		name  string
		input []int32
		d     int32
		want  []int32
	}{
		{
			name:  "test case #1",
			input: []int32{1, 2, 3, 4, 5},
			d:     4,
			want:  []int32{5, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotLeft(tt.input, tt.d)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
