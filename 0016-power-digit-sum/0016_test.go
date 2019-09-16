package main

import (
	"math/big"
	"testing"
)

func TestSumIntDigits(t *testing.T) {
	got := SumIntDigits(32768)
	want := 26

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestTwoPower(t *testing.T) {
	got := TwoPower(15)
	want := big.NewInt(32768)

	if got.Cmp(want) != 0 {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
