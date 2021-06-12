package main

import (
	"fmt"
	"testing"
)

func TestTruncate(t *testing.T) {
	examples := []struct {
		num    int
		expect int
	}{
		{num: 1, expect: 1},
		{num: 20, expect: 20},
		{num: 99991234567890, expect: 1234567890},
		{num: 999999991234567890, expect: 1234567890},
		{num: 999999991234567890, expect: 1234567890},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := Truncate(tt.num)
			want := tt.expect

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestExpand(t *testing.T) {
	examples := []struct {
		num    int
		expect int
	}{
		{num: 1, expect: 1},
		{num: 2, expect: 4},
		{num: 15, expect: 380859375},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := Expand(tt.num)
			want := tt.expect

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestFindSeries(t *testing.T) {
	examples := []struct {
		num    int
		expect int
	}{
		{num: 1, expect: 1},
		{num: 2, expect: 5},
		{num: 10, expect: 405071317},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := FindSeries(tt.num)
			want := tt.expect

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}
