package hackerrank

import "testing"

func TestWhatFlavors(t *testing.T) {
	tests := []struct {
		name  string
		cost  []int32
		money int32
		want  string
	}{
		{
			name:  "test case #1",
			cost:  []int32{2, 1, 3, 5, 6},
			money: 5,
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			whatFlavors(tt.cost, tt.money)
		})
	}
}
