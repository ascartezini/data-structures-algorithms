package hackerrank

import "testing"

func TestMakeAnagram(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want int32
	}{
		{
			name: "test case #1",
			a:    "cde",
			b:    "dcf",
			want: 2,
		},
		{
			name: "test case #2",
			a:    "fcrxzwscanmligyxyvym",
			b:    "jxwtrhvujlmrpdoqbisbwhmgpmeoke",
			want: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeAnagram(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
