package main

import "testing"

func TestFindPaths(t *testing.T) {
	t.Run("Find paths for 2x2 grid", func(t *testing.T) {
		got := FindPaths(2)
		want := 6

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
	t.Run("Find paths for 3x3 grid", func(t *testing.T) {
		got := FindPaths(3)
		want := 20

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
}
