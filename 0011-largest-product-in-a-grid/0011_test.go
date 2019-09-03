package main

import "testing"

var TestGrid = [][]int{
	{26, 38, 40, 67},
	{95, 63, 94, 39},
	{97, 17, 78, 78},
	{20, 45, 35, 14},
}

func TestGridSumming(t *testing.T) {
	sumVals := []struct {
		name                   string
		fn                     func(int, int, int, *[][]int) (int, error)
		x, y, sumCount, expect int
	}{
		{name: "Horizontal Sum 1", fn: SumHorizontal, x: 3, y: 0, sumCount: 1, expect: 67},
		{name: "Horizontal Sum 2", fn: SumHorizontal, x: 2, y: 1, sumCount: 2, expect: 133},
		{name: "Horizontal Sum 3", fn: SumHorizontal, x: 1, y: 2, sumCount: 3, expect: 173},
		{name: "Horizontal Sum 4", fn: SumHorizontal, x: 0, y: 3, sumCount: 4, expect: 114},
		{name: "Vertical Sum 1", fn: SumVertical, x: 0, y: 3, sumCount: 1, expect: 20},
		{name: "Vertical Sum 2", fn: SumVertical, x: 1, y: 2, sumCount: 2, expect: 62},
		{name: "Vertical Sum 3", fn: SumVertical, x: 2, y: 1, sumCount: 3, expect: 207},
		{name: "Vertical Sum 4", fn: SumVertical, x: 3, y: 0, sumCount: 4, expect: 198},
		{name: "Diagonal Right Sum 1", fn: SumDiagRight, x: 3, y: 3, sumCount: 1, expect: 14},
		{name: "Diagonal Right Sum 2", fn: SumDiagRight, x: 2, y: 2, sumCount: 2, expect: 92},
		{name: "Diagonal Right Sum 3", fn: SumDiagRight, x: 1, y: 1, sumCount: 3, expect: 155},
		{name: "Diagonal Right Sum 4", fn: SumDiagRight, x: 0, y: 0, sumCount: 4, expect: 181},
		{name: "Diagonal Left Sum 1", fn: SumDiagLeft, x: 0, y: 3, sumCount: 1, expect: 20},
		{name: "Diagonal Left Sum 2", fn: SumDiagLeft, x: 1, y: 2, sumCount: 2, expect: 37},
		{name: "Diagonal Left Sum 3", fn: SumDiagLeft, x: 2, y: 1, sumCount: 3, expect: 131},
		{name: "Diagonal Left Sum 4", fn: SumDiagLeft, x: 3, y: 0, sumCount: 4, expect: 198},
	}

	for _, tt := range sumVals {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fn(tt.x, tt.y, tt.sumCount, &TestGrid)
			want := tt.expect

			if err != nil {
				t.Fatalf("Got an error %q and didn't expect one", err)
			}

			if got != want {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestGridSearchOutOfBounds(t *testing.T) {
	sumVals := []struct {
		name           string
		fn             func(int, int, int, *[][]int) (int, error)
		x, y, sumCount int
	}{
		{name: "Horizontal OOB 1", fn: SumHorizontal, x: 3, y: 0, sumCount: 2},
		{name: "Horizontal OOB 2", fn: SumHorizontal, x: -1, y: 0, sumCount: 2},
		{name: "Horizontal OOB 3", fn: SumHorizontal, x: 1, y: -1, sumCount: 2},
		{name: "Vertical OOB 1", fn: SumVertical, x: 3, y: 3, sumCount: 2},
		{name: "Vertical OOB 2", fn: SumVertical, x: -1, y: 0, sumCount: 2},
		{name: "Vertical OOB 3", fn: SumVertical, x: 1, y: -1, sumCount: 2},
		{name: "DiagonalRight OOB 1", fn: SumDiagRight, x: 3, y: 3, sumCount: 2},
		{name: "DiagonalRight OOB 2", fn: SumDiagRight, x: 2, y: 3, sumCount: 2},
		{name: "DiagonalRight OOB 3", fn: SumDiagRight, x: 3, y: 2, sumCount: 2},
		{name: "DiagonalRight OOB 4", fn: SumDiagRight, x: 0, y: -1, sumCount: 2},
		{name: "DiagonalRight OOB 5", fn: SumDiagRight, x: 1, y: -1, sumCount: 2},
		{name: "DiagonalLeft OOB 1", fn: SumDiagLeft, x: 0, y: 3, sumCount: 2},
		{name: "DiagonalLeft OOB 2", fn: SumDiagLeft, x: 1, y: 3, sumCount: 2},
		{name: "DiagonalLeft OOB 3", fn: SumDiagLeft, x: 0, y: 2, sumCount: 2},
		{name: "DiagonalLeft OOB 4", fn: SumDiagLeft, x: 0, y: -1, sumCount: 2},
		{name: "DiagonalLeft OOB 5", fn: SumDiagLeft, x: 1, y: -1, sumCount: 2},
	}
	for _, tt := range sumVals {
		t.Run(tt.name, func(t *testing.T) {
			_, got := tt.fn(tt.x, tt.y, tt.sumCount, &TestGrid)
			want := GridIndexError

			if got != want {
				t.Errorf("Got error %q, wanted %q", got, want)
			}
		})
	}
}
