package main

import "testing"

func TestFindMultiples(t *testing.T) {
	t.Run("Find multiples under 10", func(t *testing.T) {
		want := 23
		got := FindMultiples(10)

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
}
