package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	examples := []struct {
		num    int
		expect []int
	}{
		{num: 14, expect: []int{2, 7}},
		{num: 15, expect: []int{3, 5}},
		{num: 644, expect: []int{2, 2, 7, 23}},
		{num: 645, expect: []int{3, 5, 43}},
		{num: 646, expect: []int{2, 17, 19}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := PrimeFactors(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestUniquePrimeFactors(t *testing.T) {
	examples := []struct {
		num    int
		expect []int
	}{
		{num: 14, expect: []int{2, 7}},
		{num: 15, expect: []int{3, 5}},
		{num: 644, expect: []int{2, 7, 23}},
		{num: 645, expect: []int{3, 5, 43}},
		{num: 646, expect: []int{2, 17, 19}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := UniquePrimeFactors(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestFindFirstSequential(t *testing.T) {
	examples := []struct {
		num    int
		expect []int
	}{
		{num: 2, expect: []int{14, 15}},
		{num: 3, expect: []int{644, 645, 646}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := FindFirstSequential(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}
