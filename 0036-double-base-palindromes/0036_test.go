package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToBinary(t *testing.T) {
	examples := []struct {
		num    int
		expect string
	}{
		{num: 0, expect: ""},
		{num: 1, expect: "1"},
		{num: 2, expect: "10"},
		{num: 585, expect: "1001001001"},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := ToBinary(tt.num)
			want := tt.expect

			if got != want {
				t.Errorf("Got %s, wanted %s", got, want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	examples := []struct {
		num    string
		expect bool
	}{
		{num: "1", expect: true},
		{num: "10", expect: false},
		{num: "11", expect: true},
		{num: "111", expect: true},
		{num: "110", expect: false},
		{num: "1101", expect: false},
		{num: "1001", expect: true},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %s", tt.num), func(t *testing.T) {
			got := IsPalindrome(tt.num)
			want := tt.expect

			if got != want {
				t.Errorf("Got %t, wanted %t", got, want)
			}
		})
	}
}

func TestIsDoubleBasePalindrome(t *testing.T) {
	examples := []struct {
		num    int
		expect bool
	}{
		{num: 1, expect: true},
		{num: 2, expect: false},
		{num: 3, expect: true},
		{num: 4, expect: false},
		{num: 5, expect: true},
		{num: 585, expect: true},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := IsDoubleBasePalindrome(tt.num)
			want := tt.expect

			if got != want {
				t.Errorf("Got %t, wanted %t", got, want)
			}
		})
	}
}

func TestFindDBPalindromesUnder(t *testing.T) {
	examples := []struct {
		num    int
		expect []int
	}{
		{num: 10, expect: []int{1, 3, 5, 7, 9}},
		{num: 1000, expect: []int{1, 3, 5, 7, 9, 33, 99, 313, 585, 717}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := FindDBPalindromesUnder(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestSumDBPalindromesUnder(t *testing.T) {
	examples := []struct {
		num    int
		expect int
	}{
		{num: 10, expect: 25},
		{num: 100, expect: 157},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := SumDBPalindromesUnder(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}
