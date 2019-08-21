package main

import "testing"

func TestGetSumOfSquares(t *testing.T) {
	t.Run("Get sum of squares", func(t *testing.T) {
		got := GetSumOfSquares(10)
		want := 385

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
}

func TestSquareOfSum(t *testing.T) {
	t.Run("Get square of sum", func(t *testing.T) {
		got := GetSquareOfSum(10)
		want := 3025

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
}

func TestGetDifference(t *testing.T) {
	t.Run("Get the difference", func(t *testing.T) {
		got := GetDifference(10)
		want := 2640

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
}
