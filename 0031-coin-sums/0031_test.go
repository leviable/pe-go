package main

import (
	"fmt"
	"testing"
)

func TestFindOnePence(t *testing.T) {
	got := FindOnePence(0)
	want := 1

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestFindTwoPence(t *testing.T) {
	evalNums := []struct {
		total, expect int
	}{
		{total: 198, expect: 2},
		{total: 199, expect: 1},
		{total: 200, expect: 1},
	}
	for _, tt := range evalNums {
		t.Run(fmt.Sprintf("Evaluating %d", tt.total), func(t *testing.T) {
			got := FindTwoPence(tt.total)
			want := tt.expect

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestFindFivePence(t *testing.T) {
	evalNums := []struct {
		total, expect int
	}{
		{total: 190, expect: 10},
		{total: 200, expect: 1},
	}
	for _, tt := range evalNums {
		t.Run(fmt.Sprintf("Evaluating %d", tt.total), func(t *testing.T) {
			got := FindFivePence(tt.total)
			want := tt.expect

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}
