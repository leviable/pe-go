package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNextCollatzNum(t *testing.T) {
	TestTerms := []struct {
		value, expect int
	}{
		{value: 13, expect: 40},
		{value: 40, expect: 20},
		{value: 20, expect: 10},
		{value: 10, expect: 5},
		{value: 5, expect: 16},
		{value: 16, expect: 8},
		{value: 8, expect: 4},
		{value: 4, expect: 2},
		{value: 2, expect: 1},
	}

	for _, tt := range TestTerms {
		t.Run(fmt.Sprintf("Testing %d", tt.value), func(t *testing.T) {
			got := NextCollatzNum(tt.value)
			want := tt.expect

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestFindCollatzSeq(t *testing.T) {
	TestTerms := []struct {
		value  int
		expect []int
	}{
		{value: 13, expect: []int{1, 2, 4, 8, 16, 5, 10, 20, 40}},
		{value: 40, expect: []int{1, 2, 4, 8, 16, 5, 10, 20}},
		{value: 20, expect: []int{1, 2, 4, 8, 16, 5, 10}},
		{value: 10, expect: []int{1, 2, 4, 8, 16, 5}},
		{value: 5, expect: []int{1, 2, 4, 8, 16}},
		{value: 16, expect: []int{1, 2, 4, 8}},
		{value: 8, expect: []int{1, 2, 4}},
		{value: 4, expect: []int{1, 2}},
		{value: 2, expect: []int{1}},
	}
	for _, tt := range TestTerms {
		t.Run(fmt.Sprintf("Getting sequence for %d", tt.value), func(t *testing.T) {
			got := GetCollatzSeq(tt.value)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestGetLongestCollatzSeq(t *testing.T) {
	got := GetLongestCollatzSeq(13)
	want := LongestSequence{9, 20}

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
