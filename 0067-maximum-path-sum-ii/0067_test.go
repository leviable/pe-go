package main

import (
	"fmt"
	"testing"
)

var TestTriangle = [][]int{
	{3},
	{7, 4},
	{2, 4, 6},
	{8, 5, 9, 3},
}

func TestReducePaths(t *testing.T) {
	testTerms := []struct {
		x, y, expected int
	}{
		{x: 0, y: 3, expected: 20},
		{x: 1, y: 3, expected: 19},
		{x: 2, y: 3, expected: 23},
		{x: 3, y: 3, expected: 16},
	}

	for _, tt := range testTerms {
		t.Run(fmt.Sprintf("Evaluating (%d, %d)", tt.x, tt.y), func(t *testing.T) {
			got := ReducePaths(TestTriangle, tt.x, tt.y)
			want := tt.expected

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestFindGreatestPath(t *testing.T) {
	got := FindGreatestPath(TestTriangle)
	want := 23

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
