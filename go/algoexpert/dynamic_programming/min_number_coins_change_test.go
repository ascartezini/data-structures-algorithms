package algoexpert

import (
	"testing"
)

func TestMinNumberOfCoinsForChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "test case #1",
			coins:  []int{1, 2, 4},
			amount: 6,
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minNumberOfCoinsForChange(tt.amount, tt.coins)

			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
