package main

import "testing"

func TestGetLargestPrimeFactor(t *testing.T) {
	t.Run("get largest prime factor", func(t *testing.T) {
		got := GetLargestPrimeFactor(13195)
		want := 29

		if got != want {
			t.Errorf("Wanted %d, got %d", want, got)
		}
	})
}
