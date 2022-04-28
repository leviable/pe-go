package main

import (
	"reflect"
	"testing"
)

func TestConvertNumber(t *testing.T) {
	got := ConvertNumber(1_000_000)
	want := Num{num: 1.0, exp: 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestGetFactiorial(t *testing.T) {
	got := GetFactorial(5)
	want := []Num{
		Num{5, 0},
		Num{4, 0},
		Num{3, 0},
		Num{2, 0},
		Num{1, 0},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestCompute(t *testing.T) {
	got := Compute(5, 3)
	want := Num{10, 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestNormalize(t *testing.T) {
	got := Normalize(Num{10, 0})
	want := Num{1, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestIsOver(t *testing.T) {
	got := IsOver(Num{1.0, 4}, 3)
	want := true

	if got != want {
		t.Errorf("Got %t, want %t", got, want)
	}
}

func TestPE0053(t *testing.T) {
	got := PE0053(23, 6)
	want := 4

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
