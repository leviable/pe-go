package main

import (
	"testing"
)

func TestProductIsPalindrome(t *testing.T) {
	t.Run("is palindrome", func(t *testing.T) {
		num := 9009
		got := IsPalindrome(num)
		want := true

		if got != want {
			t.Errorf("For number %d: wanted %t, got %t", num, want, got)
		}

	})
	t.Run("is not palindrome", func(t *testing.T) {
		num := 9090
		got := IsPalindrome(num)
		want := false

		if got != want {
			t.Errorf("For number %d: wanted %t, got %t", num, want, got)
		}

	})
}

func TestGetLargestPalindrome(t *testing.T) {
	t.Run("largest for 2 digits", func(t *testing.T) {
		got := GetLargestPalindrome(2)
		want := 9009

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
}
