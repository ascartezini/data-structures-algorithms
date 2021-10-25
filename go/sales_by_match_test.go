package main

import (
	"testing"
)

func TestSalesByMatch(t *testing.T) {
	tests := []struct {
		name  string
		input []int32
		k     int32
		want  int32
	}{
		{
			name:  "",
			input: []int32{1, 2, 1, 2, 1, 3, 2},
			k:     7,
			want:  2,
		},
		{
			name:  "",
			input: []int32{10, 20, 20, 10, 10, 30, 50, 10, 20},
			k:     9,
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sockMerchant(tt.k, tt.input)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
