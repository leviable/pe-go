package main

import "testing"

var TestGrid = [][]int{
	{26, 38, 40, 67},
	{95, 63, 94, 39},
	{97, 17, 78, 78},
	{20, 45, 35, 14},
}

func TestGridSumming(t *testing.T) {
	t.Run("Sum horizontal", func(t *testing.T) {
		got := SumHorizontal(0, 0, 2, &TestGrid)
		want := 64

		if got != want {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})
}
