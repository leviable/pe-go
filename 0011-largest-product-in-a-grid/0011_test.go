package main

import "testing"

var TestGrid = [][]int{
	{26, 38, 40, 67},
	{95, 63, 94, 39},
	{97, 17, 78, 78},
	{20, 45, 35, 14},
}

func TestGridSumming(t *testing.T) {
	ProdVals := []struct {
		name                    string
		fn                      func(int, int, int, *[][]int) (int, error)
		x, y, prodCount, expect int
	}{
		{name: "Horizontal prod 1", fn: ProdHorizontal, x: 3, y: 0, prodCount: 1, expect: 67},
		{name: "Horizontal prod 2", fn: ProdHorizontal, x: 2, y: 1, prodCount: 2, expect: 3666},
		{name: "Horizontal prod 3", fn: ProdHorizontal, x: 1, y: 2, prodCount: 3, expect: 103428},
		{name: "Horizontal prod 4", fn: ProdHorizontal, x: 0, y: 3, prodCount: 4, expect: 441000},
		{name: "Vertical prod 1", fn: ProdVertical, x: 0, y: 3, prodCount: 1, expect: 20},
		{name: "Vertical prod 2", fn: ProdVertical, x: 1, y: 2, prodCount: 2, expect: 765},
		{name: "Vertical prod 3", fn: ProdVertical, x: 2, y: 1, prodCount: 3, expect: 256620},
		{name: "Vertical prod 4", fn: ProdVertical, x: 3, y: 0, prodCount: 4, expect: 2853396},
		{name: "Diagonal Right prod 1", fn: ProdDiagRight, x: 3, y: 3, prodCount: 1, expect: 14},
		{name: "Diagonal Right prod 2", fn: ProdDiagRight, x: 2, y: 2, prodCount: 2, expect: 1092},
		{name: "Diagonal Right prod 3", fn: ProdDiagRight, x: 1, y: 1, prodCount: 3, expect: 68796},
		{name: "Diagonal Right prod 4", fn: ProdDiagRight, x: 0, y: 0, prodCount: 4, expect: 1788696},
		{name: "Diagonal Left prod 1", fn: ProdDiagLeft, x: 0, y: 3, prodCount: 1, expect: 20},
		{name: "Diagonal Left prod 2", fn: ProdDiagLeft, x: 1, y: 2, prodCount: 2, expect: 340},
		{name: "Diagonal Left prod 3", fn: ProdDiagLeft, x: 2, y: 1, prodCount: 3, expect: 31960},
		{name: "Diagonal Left prod 4", fn: ProdDiagLeft, x: 3, y: 0, prodCount: 4, expect: 2141320},
	}

	for _, tt := range ProdVals {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fn(tt.x, tt.y, tt.prodCount, &TestGrid)
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
	ProdVals := []struct {
		name            string
		fn              func(int, int, int, *[][]int) (int, error)
		x, y, prodCount int
	}{
		{name: "Horizontal OOB 1", fn: ProdHorizontal, x: 3, y: 0, prodCount: 2},
		{name: "Horizontal OOB 2", fn: ProdHorizontal, x: -1, y: 0, prodCount: 2},
		{name: "Horizontal OOB 3", fn: ProdHorizontal, x: 1, y: -1, prodCount: 2},
		{name: "Vertical OOB 1", fn: ProdVertical, x: 3, y: 3, prodCount: 2},
		{name: "Vertical OOB 2", fn: ProdVertical, x: -1, y: 0, prodCount: 2},
		{name: "Vertical OOB 3", fn: ProdVertical, x: 1, y: -1, prodCount: 2},
		{name: "DiagonalRight OOB 1", fn: ProdDiagRight, x: 3, y: 3, prodCount: 2},
		{name: "DiagonalRight OOB 2", fn: ProdDiagRight, x: 2, y: 3, prodCount: 2},
		{name: "DiagonalRight OOB 3", fn: ProdDiagRight, x: 3, y: 2, prodCount: 2},
		{name: "DiagonalRight OOB 4", fn: ProdDiagRight, x: 0, y: -1, prodCount: 2},
		{name: "DiagonalRight OOB 5", fn: ProdDiagRight, x: 1, y: -1, prodCount: 2},
		{name: "DiagonalLeft OOB 1", fn: ProdDiagLeft, x: 0, y: 3, prodCount: 2},
		{name: "DiagonalLeft OOB 2", fn: ProdDiagLeft, x: 1, y: 3, prodCount: 2},
		{name: "DiagonalLeft OOB 3", fn: ProdDiagLeft, x: 0, y: 2, prodCount: 2},
		{name: "DiagonalLeft OOB 4", fn: ProdDiagLeft, x: 0, y: -1, prodCount: 2},
		{name: "DiagonalLeft OOB 5", fn: ProdDiagLeft, x: 1, y: -1, prodCount: 2},
	}
	for _, tt := range ProdVals {
		t.Run(tt.name, func(t *testing.T) {
			_, got := tt.fn(tt.x, tt.y, tt.prodCount, &TestGrid)
			want := GridIndexError

			if got != want {
				t.Errorf("Got error %q, wanted %q", got, want)
			}
		})
	}
}

func TestGetMaxProduct(t *testing.T) {
	t.Run("Max product for test grid", func(t *testing.T) {
		got, _ := GetMaxProduct(4, &TestGrid)
		want := 21941010

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
}
