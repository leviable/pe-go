package main

import "testing"

func TestIsTriplet(t *testing.T) {
	t.Run("Is triplet 3-4-5", func(t *testing.T) {
		got := IsTriplet(3, 4, 5)
		want := true

		if got != want {
			t.Errorf("Got %t, wanted %t", got, want)
		}
	})
	t.Run("Is triplet 5-12-13", func(t *testing.T) {
		got := IsTriplet(5, 12, 13)
		want := true

		if got != want {
			t.Errorf("Got %t, wanted %t", got, want)
		}
	})
	t.Run("Is not triplet 4-3-5", func(t *testing.T) {
		got := IsTriplet(4, 3, 5)
		want := false

		if got != want {
			t.Errorf("Got %t, wanted %t", got, want)
		}
	})
	t.Run("Is not triplet 3-5-4", func(t *testing.T) {
		got := IsTriplet(3, 5, 4)
		want := false

		if got != want {
			t.Errorf("Got %t, wanted %t", got, want)
		}
	})
}
