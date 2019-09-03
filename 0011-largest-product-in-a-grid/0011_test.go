package main

import "testing"

var TestGrid = [][]int{
	{26, 38, 40, 67},
	{95, 63, 94, 39},
	{97, 17, 78, 78},
	{20, 45, 35, 14},
}

func TestHorizontalGridSumming(t *testing.T) {
	sumHoriz := []struct {
		name                   string
		x, y, sumCount, expect int
	}{
		{name: "Horizontal Sum 1", x: 3, y: 0, sumCount: 1, expect: 67},
		{name: "Horizontal Sum 2", x: 2, y: 1, sumCount: 2, expect: 133},
		{name: "Horizontal Sum 3", x: 1, y: 2, sumCount: 3, expect: 173},
		{name: "Horizontal Sum 4", x: 0, y: 3, sumCount: 4, expect: 114},
	}

	for _, tt := range sumHoriz {
		t.Run(tt.name, func(t *testing.T) {
			got := SumHorizontal(tt.x, tt.y, tt.sumCount, &TestGrid)
			want := tt.expect

			if got != want {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestVerticalGridSumming(t *testing.T) {
	sumHoriz := []struct {
		name                   string
		x, y, sumCount, expect int
	}{
		{name: "Vertical Sum 1", x: 0, y: 3, sumCount: 1, expect: 20},
		{name: "Vertical Sum 2", x: 1, y: 2, sumCount: 2, expect: 62},
		{name: "Vertical Sum 3", x: 2, y: 1, sumCount: 3, expect: 207},
		{name: "Vertical Sum 4", x: 3, y: 0, sumCount: 4, expect: 198},
	}

	for _, tt := range sumHoriz {
		t.Run(tt.name, func(t *testing.T) {
			got := SumVertical(tt.x, tt.y, tt.sumCount, &TestGrid)
			want := tt.expect

			if got != want {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestDiagonalRightGridSumming(t *testing.T) {
	sumHoriz := []struct {
		name                   string
		x, y, sumCount, expect int
	}{
		{name: "Diagonal Right Sum 1", x: 3, y: 3, sumCount: 1, expect: 14},
		{name: "Diagonal Right Sum 2", x: 2, y: 2, sumCount: 2, expect: 92},
		{name: "Diagonal Right Sum 3", x: 1, y: 1, sumCount: 3, expect: 155},
		{name: "Diagonal Right Sum 4", x: 0, y: 0, sumCount: 4, expect: 181},
	}

	for _, tt := range sumHoriz {
		t.Run(tt.name, func(t *testing.T) {
			got := SumDiagonalRight(tt.x, tt.y, tt.sumCount, &TestGrid)
			want := tt.expect

			if got != want {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestDiagonalLeftGridSumming(t *testing.T) {
	sumHoriz := []struct {
		name                   string
		x, y, sumCount, expect int
	}{
		{name: "Diagonal Left Sum 1", x: 0, y: 3, sumCount: 1, expect: 20},
		{name: "Diagonal Left Sum 2", x: 1, y: 2, sumCount: 2, expect: 37},
		{name: "Diagonal Left Sum 3", x: 2, y: 1, sumCount: 3, expect: 131},
		{name: "Diagonal Left Sum 4", x: 3, y: 0, sumCount: 4, expect: 198},
	}

	for _, tt := range sumHoriz {
		t.Run(tt.name, func(t *testing.T) {
			got := SumDiagonalLeft(tt.x, tt.y, tt.sumCount, &TestGrid)
			want := tt.expect

			if got != want {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}
