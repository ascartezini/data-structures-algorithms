package hackerrank

import "testing"

func TestMaximumSum(t *testing.T) {
	tests := []struct {
		name string
		a    []int64
		m    int64
		want int64
	}{
		{
			name: "test case #1",
			a:    []int64{-2, 2, 5, -11, 6},
			m:    5,
			want: 7,
		},
		{
			name: "test case #2",
			a:    []int64{-2, 3, 2, -1},
			m:    5,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumSum(tt.a)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
