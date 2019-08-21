package main

import (
	"testing"
)

func TestGetSmallestMultiple(t *testing.T) {
	t.Run("smallest multiple from 1 to 10", func(t *testing.T) {
		got := GetSmallestMultiple(10)
		want := 2520

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
}
