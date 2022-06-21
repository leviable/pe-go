package main

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestExpand(t *testing.T) {
	examples := []struct {
		expand   int
		num, den *big.Int
	}{
		{1, big.NewInt(3), big.NewInt(2)},
		{2, big.NewInt(7), big.NewInt(5)},
		{3, big.NewInt(17), big.NewInt(12)},
		{4, big.NewInt(41), big.NewInt(29)},
		{5, big.NewInt(99), big.NewInt(70)},
		{6, big.NewInt(239), big.NewInt(169)},
		{7, big.NewInt(577), big.NewInt(408)},
		{8, big.NewInt(1393), big.NewInt(985)},
	}

	for _, tt := range examples {

		t.Run(fmt.Sprintf("Expansion %d", tt.expand), func(t *testing.T) {
			got := make([]*big.Int, 2)
			got[0], got[1] = Expand(tt.expand)
			want := []*big.Int{tt.num, tt.den}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestIsBigger(t *testing.T) {
	examples := []struct {
		expand   int
		isBigger bool
	}{
		{1, false},
		{2, false},
		{3, false},
		{4, false},
		{5, false},
		{6, false},
		{7, false},
		{8, true},
	}

	for _, tt := range examples {

		t.Run(fmt.Sprintf("Expansion %d", tt.expand), func(t *testing.T) {
			num, den := Expand(tt.expand)
			got := IsBigger(num, den)
			want := tt.isBigger

			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}
		})
	}
}
