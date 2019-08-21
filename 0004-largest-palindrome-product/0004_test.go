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
