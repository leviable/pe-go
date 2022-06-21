package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMult(t *testing.T) {
	examples := []struct {
		num, pow int
		expect   []int
	}{
		{2, 2, []int{4}},
		{2, 3, []int{8}},
		{2, 4, []int{1, 6}},
		{2, 5, []int{3, 2}},
		{2, 7, []int{1, 2, 8}},
	}

	for _, tt := range examples {
		got := Mult(tt.num, tt.pow)
		want := tt.expect

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestSumArray(t *testing.T) {
	examples := []struct {
		num    []int
		expect int
	}{
		{[]int{1, 2, 8}, 11},
		{Mult(10, 100), 1},
		{Mult(100, 10), 1},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Checking %d", tt.num), func(t *testing.T) {
			got := SumArray(tt.num)
			want := tt.expect

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
