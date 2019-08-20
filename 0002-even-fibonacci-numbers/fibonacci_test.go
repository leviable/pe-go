package main

import "testing"

func TestSumEvenValuedTerms(t *testing.T) {
	t.Run("sum even valued terms of fib sequence", func(t *testing.T) {
		want := 44
		got := SumEvenValuedTerms(100)

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
}
