package main

import (
	"testing"
)

func TestVowelSubstring(t *testing.T) {
	tests := []struct {
		name  string
		input string
		k     int32
		want  string
	}{
		{
			name:  "",
			input: "azerdii",
			k:     5,
			want:  "erdii",
		},
		{
			name:  "",
			input: "fsiauihfiweifasnfashidufasfuiashufiahsufihsauhfduasda",
			k:     5,
			want:  "siaui",
		},
		{
			name:  "",
			input: "andregregorioscartezini",
			k:     5,
			want:  "egori",
		},
		{
			name:  "",
			input: "wtlsk",
			k:     5,
			want:  "Not found!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubstring(tt.input, 5)

			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
