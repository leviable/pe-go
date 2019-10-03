package main

import (
	"fmt"
	"testing"
)

func TestGetPowerSum(t *testing.T) {
	evalNums := []struct {
		num, pow, expect int
	}{
		{num: 1634, pow: 4, expect: 1634},
		{num: 8208, pow: 4, expect: 8208},
		{num: 9474, pow: 4, expect: 9474},
	}

	for _, tt := range evalNums {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := GetPowerSum(tt.num, tt.pow)
			want := tt.expect

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestSumAllForPower(t *testing.T) {
	got := SumAllForPower(4)
	want := 19316

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
