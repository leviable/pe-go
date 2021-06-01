package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestGeneratePrimes(t *testing.T) {
	examples := []struct {
		limit  int
		expect []int
	}{
		{limit: 4, expect: []int{2, 3, 5, 7}},
		{limit: 5, expect: []int{2, 3, 5, 7, 11}},
		{limit: 6, expect: []int{2, 3, 5, 7, 11, 13}},
		{limit: 7, expect: []int{2, 3, 5, 7, 11, 13, 17}},
		{limit: 8, expect: []int{2, 3, 5, 7, 11, 13, 17, 19}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.limit), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			got := make([]int, 0)

		loop:
			for p := range GeneratePrimes(ctx) {
				got = append(got, p)
				if len(got) == tt.limit {
					cancel()
					break loop
				}
			}
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	examples := []struct {
		num       int
		direction string
		expect    []int
	}{
		{num: 3797, direction: RIGHT, expect: []int{3797, 379, 37, 3}},
		{num: 3797, direction: LEFT, expect: []int{3797, 797, 97, 7}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := make([]int, 0)
			for n := range Truncate(tt.num, tt.direction) {
				got = append(got, n)
			}
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestFindTruncatablePrimes(t *testing.T) {
	examples := []struct {
		num    int
		expect []int
	}{
		{num: 1, expect: []int{23}},
		{num: 2, expect: []int{23, 37}},
		{num: 3, expect: []int{23, 37, 53}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			got := make([]int, 0)
		loop:
			for p := range FindTruncatablePrimes(ctx) {
				got = append(got, p)
				if len(got) >= tt.num {
					cancel()
					break loop
				}
			}
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}
