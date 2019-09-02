package main

import "reflect"
import "testing"

func TestSieve(t *testing.T) {
	t.Run("Testing sieve generation - 20", func(t *testing.T) {
		got := GetPrimesUnder(20)
		want := []int{2, 3, 5, 7, 11, 13, 17, 19}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})
}

func TestSumPrimes(t *testing.T) {
	t.Run("Testing Sum Primes", func(t *testing.T) {
		got := SumPrimes(&[]int{2, 3, 5, 7})
		want := 17

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
}

func BenchmarkSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimesUnder(20)
	}
}
