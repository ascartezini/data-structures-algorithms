package hackerrank

import (
	"testing"
)

func TestCountingValleys(t *testing.T) {
	tests := []struct {
		name  string
		steps int32
		path  string
		want  int32
	}{
		{
			name:  "",
			steps: 8,
			path:  "UDDDUDUU",
			want:  1,
		},
		{
			name:  "",
			steps: 12,
			path:  "DDUUDDUDUUUD",
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countingValleys(tt.steps, tt.path)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
