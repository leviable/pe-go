package main

import "fmt"
import "testing"

func TestIsPrime(t *testing.T) {

	primeCandidates := []struct {
		value  int
		expect bool
	}{
		{value: 1, expect: false},
		{value: 2, expect: true},
		{value: 3, expect: true},
		{value: 4, expect: false},
		{value: 5, expect: true},
		{value: 6, expect: false},
		{value: 7, expect: true},
		{value: 8, expect: false},
		{value: 9, expect: false},
		{value: 10, expect: false},
		{value: 11, expect: true},
	}

	for _, tt := range primeCandidates {
		t.Run(fmt.Sprintf("Evaluating %d", tt.value), func(t *testing.T) {
			got := IsPrime(tt.value)
			want := tt.expect
			if got != want {
				t.Errorf("For %d, got %t, want %t", tt.value, got, want)
			}
		})
	}
}

func TestFindNthPrime(t *testing.T) {
	t.Run("(Brute Force)Find the 6th prime", func(t *testing.T) {
		got := FindNthPrime(6)
		want := 13

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
	t.Run("(Parallel) Find the 6th prime", func(t *testing.T) {
		got := FindNthPrimePar(6)
		want := 13

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
}
