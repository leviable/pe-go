package main

import (
	"fmt"
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	testVals := []struct {
		num, expected int
	}{
		{num: 1, expected: 1},
		{num: 2, expected: 2},
		{num: 3, expected: 6},
		{num: 4, expected: 24},
		{num: 5, expected: 120},
	}

	for _, tt := range testVals {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := Factorial(tt.num)
			want := big.NewInt(int64(tt.expected))

			if got.Cmp(want) != 0 {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestSumDigits(t *testing.T) {
	got := SumDigits(3628800)
	want := 27

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestFindFactorialSum(t *testing.T) {
	got := FindFactorialSum(10)
	want := 27

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
