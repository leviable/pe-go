package main

import "testing"

func TestFindLargestProduct(t *testing.T) {
	t.Run("Find largest product of 4 adjacent", func(t *testing.T) {
		got := FindLargestProduct(4)
		want := 5832

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
}

func TestGetProduct(t *testing.T) {
	t.Run("find product of 1234", func(t *testing.T) {
		got := GetProduct("1234")
		want := 24

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
}
