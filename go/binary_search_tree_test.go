package main

import (
	"testing"
)

func TestBinarySearchThree(t *testing.T) {

	bst := &BST{
		val: 20,
		left: &BST{val: 10,
			left: &BST{val: 8,
				left: &BST{val: 6}},
			right: &BST{val: 12,
				left:  &BST{val: 11},
				right: &BST{val: 13}}},
		right: &BST{val: 30,
			left: &BST{val: 25,
				left: &BST{val: 23}},
			right: &BST{val: 40}}}

	tests := []struct {
		name   string
		input  *BST
		target int
		want   bool
	}{
		{
			name:   "test case #1",
			input:  bst,
			target: 25,
			want:   true,
		},
		{
			name:   "test case #2",
			input:  bst,
			target: 27,
			want:   false,
		},
		{
			name:   "test case #3",
			input:  bst,
			target: 6,
			want:   true,
		},
		{
			name:   "test case #4",
			input:  bst,
			target: 13,
			want:   true,
		},
		{
			name:   "test case #4",
			input:  bst,
			target: 41,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPresent(tt.input, tt.target)

			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}
