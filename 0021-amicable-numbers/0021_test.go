package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestGetDivisors(t *testing.T) {
	testVars := []struct {
		num    int
		expect []int
	}{
		{num: 1, expect: []int{1}},
		{num: 2, expect: []int{1}},
		{num: 3, expect: []int{1}},
		{num: 10, expect: []int{1, 2, 5}},
		{num: 25, expect: []int{1, 5}},
		{num: 220, expect: []int{1, 2, 4, 5, 10, 11, 20, 22, 44, 55, 110}},
		{num: 284, expect: []int{1, 2, 4, 71, 142}},
	}
	for _, tt := range testVars {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := GetDivisors(tt.num)
			want := tt.expect
			sort.Ints(got)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %+v, wanted %+v", got, want)
			}
		})
	}
}

func TestGetSum(t *testing.T) {
	testNums := []struct {
		num, expected int
	}{
		{num: 220, expected: 284},
		{num: 284, expected: 220},
	}

	for _, tt := range testNums {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := GetSum(tt.num)
			want := tt.expected

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestHasAmicable(t *testing.T) {
	testNums := []struct {
		num    int
		expect bool
	}{
		{num: 10, expect: false},
		{num: 220, expect: true},
		{num: 284, expect: true},
	}
	for _, tt := range testNums {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := HasAmicable(tt.num)
			want := tt.expect

			if got != want {
				t.Errorf("Got %t, wanted %t", got, want)
			}
		})
	}
}
