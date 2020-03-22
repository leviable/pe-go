package main

import (
	"testing"
)

func TestIsPandigital(t *testing.T) {
	t.Run("Evaluating is 1 to 9 pan digital", func(t *testing.T) {
		got := IsPandigital(39, 186, 7254)
		want := true

		if got != want {
			t.Errorf("Got %t, wanted %t", got, want)
		}
	})
	t.Run("Evaluating is not 1 to 9 pan digital", func(t *testing.T) {
		got := IsPandigital(9, 186, 7254)
		want := false

		if got != want {
			t.Errorf("Got %t, wanted %t", got, want)
		}
	})
	t.Run("Evaluating is not 1 to 9 pan digital", func(t *testing.T) {
		got := IsPandigital(999, 186, 7254)
		want := false

		if got != want {
			t.Errorf("Got %t, wanted %t", got, want)
		}
	})
}
