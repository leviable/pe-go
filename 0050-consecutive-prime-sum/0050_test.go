package main

import (
	"reflect"
	"testing"
)

func TestPrimeSieve(t *testing.T) {
	got := PrimeSieve(10)
	want := []int{2, 3, 5, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestIsPrime(t *testing.T) {
	primes := PrimeSieve(100)

	got := IsPrime(primes, 5)
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestDo0050(t *testing.T) {
	examples := []struct {
		limit, want int
	}{
		{limit: 100, want: 41},
		{limit: 1_000, want: 953},
	}
	for _, tt := range examples {
		got := Do0050(tt.limit)
		want := tt.want

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}
