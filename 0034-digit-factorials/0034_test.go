package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	examples := []struct {
		num, expect int
	}{
		{num: 1, expect: 1},
		{num: 2, expect: 2},
		{num: 3, expect: 6},
		{num: 4, expect: 24},
		{num: 5, expect: 120},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := Factorial(tt.num)
			want := tt.expect

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestSplitNum(t *testing.T) {
	examples := []struct {
		num    int
		expect []int
	}{
		{num: 1, expect: []int{1}},
		{num: 2, expect: []int{2}},
		{num: 12, expect: []int{1, 2}},
		{num: 123, expect: []int{1, 2, 3}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := SplitNum(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestIsDigitFactorial(t *testing.T) {
	examples := []struct {
		num    int
		expect bool
	}{
		{num: 145, expect: true},
		{num: 11, expect: false},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := IsDigitFactorial(tt.num)
			want := tt.expect

			if got != want {
				t.Errorf("Got %t, wanted %t", got, want)
			}
		})
	}
}
