package main

import (
	"testing"
)

func TestSumDiags(t *testing.T) {
	t.Run("Evaluating width 5", func(t *testing.T) {
		got := SumDiags(5)
		want := 101

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
	t.Run("Evaluating width 7", func(t *testing.T) {
		got := SumDiags(7)
		want := 261

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
}
