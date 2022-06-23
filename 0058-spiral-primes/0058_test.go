package main

import (
	"reflect"
	"testing"
)

func TestIsPrime(t *testing.T) {
	examples := []struct {
		num  int
		want bool
	}{
		{2, false},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
	}

	for _, tt := range examples {
		got := IsPrime(tt.num)

		if got != tt.want {
			t.Errorf("got %t, want %t", got, tt.want)
		}
	}
}

func TestRowGen(t *testing.T) {
	examples := []struct {
		row  int
		want []int
	}{
		{1, []int{1}},
		{2, []int{3, 5, 7}},
		{3, []int{13, 17, 21}},
		{4, []int{31, 37, 43}},
		{5, []int{57, 65, 73}},
		{6, []int{91, 101, 111}},
	}

	rowChan := RowGenerator()

	for _, tt := range examples {
		got := <-rowChan

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
