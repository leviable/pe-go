package main

import (
	"reflect"
	"testing"
)

func TestPrimeSieve(t *testing.T) {
	got := PrimeSieve(10)
	want := []int{2, 3, 5, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestIsPrime(t *testing.T) {
	primes := PrimeSieve(100)

	got := IsPrime(primes, 5)
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestReplace(t *testing.T) {
	primes := PrimeSieve(100)
	got := Replace(primes, 13)
	want := []int{13, 23, 43, 53, 73, 83}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestApplyMask(t *testing.T) {
	examples := []struct {
		num, replaceNum, want int
		mask                  string
		err                   error
	}{
		{num: 13, mask: "10", replaceNum: 0, want: -1, err: ErrNumStartsWithZero},
		{num: 13, mask: "10", replaceNum: 7, want: 73, err: nil},
		{num: 56113, mask: "00110", replaceNum: 0, want: 56003, err: nil},
	}
	for _, tt := range examples {
		got, err := applyMask(tt.num, tt.mask, tt.replaceNum)
		want := tt.want

		if err != tt.err {
			t.Fatalf("got error %v, want error %v", err, tt.err)
		}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}

func TestGetMasks(t *testing.T) {
	examples := []struct {
		num  int
		want []string
	}{
		{num: 13, want: []string{"10"}},
		{num: 113, want: []string{"010", "100", "110"}},
	}
	for _, tt := range examples {
		masks := make([]string, 0)
		got := getMasks(tt.num, masks, "")
		want := tt.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestExample1(t *testing.T) {
	got := PE0051(100)
	want := []int{13, 23, 43, 53, 73, 83}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestExample2(t *testing.T) {
	got := PE0051(100_000)
	want := []int{56003, 56113, 56333, 56443, 56663, 56773, 56993}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
