package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestIsAbundant(t *testing.T) {
	testNums := []struct {
		num      int
		expected bool
	}{
		{num: 10, expected: false},
		{num: 12, expected: true},
	}
	for _, tt := range testNums {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := IsAbundant(tt.num)
			want := tt.expected

			if got != want {
				t.Errorf("Got %t, wanted %t", got, want)
			}
		})
	}
}

func TestGetDivisors(t *testing.T) {
	testNums := []struct {
		num      int
		expected []int
	}{
		{num: 12, expected: []int{1, 2, 3, 4, 6}},
		{num: 25, expected: []int{1, 5}},
	}
	for _, tt := range testNums {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := GetDivisors(tt.num)
			want := tt.expected

			sort.Ints(got)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestGetAbundantNumbers(t *testing.T) {
	got := GetAbundantNumbers(13)
	want := []int{12}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %+v, wanted %+v", got, want)
	}
}
