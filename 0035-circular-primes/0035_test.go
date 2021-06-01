package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	examples := []struct {
		num    int
		expect []int
	}{
		{num: 1, expect: []int{1}},
		{num: 12, expect: []int{21, 12}},
		{num: 123, expect: []int{231, 312, 123}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := make([]int, 0)
			for v := range Rotate(tt.num) {
				got = append(got, v)
			}
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestPrimeSieve(t *testing.T) {
	examples := []struct {
		num    int
		expect []int
	}{
		{num: 10, expect: []int{2, 3, 5, 7}},
		{num: 20, expect: []int{2, 3, 5, 7, 11, 13, 17, 19}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := GetPrimesUnder(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestIsCircularPrime(t *testing.T) {
	examples := []struct {
		prime      int
		isCircular bool
	}{
		{prime: 2, isCircular: true},
		{prime: 3, isCircular: true},
		{prime: 19, isCircular: false},
		{prime: 23, isCircular: false},
	}

	primes := GetPrimesUnder(100)

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.prime), func(t *testing.T) {
			got := IsCircularPrime(tt.prime, primes)
			want := tt.isCircular

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %t, wanted %t", got, want)
			}
		})
	}
}

func TestPrunePrimeSieve(t *testing.T) {
	examples := []struct {
		num    int
		expect []int
	}{
		{num: 10, expect: []int{2, 3, 5, 7}},
		{num: 20, expect: []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{num: 30, expect: []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{num: 40, expect: []int{2, 3, 5, 7, 11, 13, 17, 19, 31, 37}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := PrunePrimes(GetPrimesUnder(tt.num))
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestGetCircularPrimesUnder(t *testing.T) {
	examples := []struct {
		num    int
		expect []int
	}{
		{num: 10, expect: []int{2, 3, 5, 7}},
		{num: 20, expect: []int{2, 3, 5, 7, 11, 13, 17}},
		{num: 30, expect: []int{2, 3, 5, 7, 11, 13, 17}},
		{num: 40, expect: []int{2, 3, 5, 7, 11, 13, 17, 31, 37}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := GetCircularPrimesUnder(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}
