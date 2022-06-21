package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsPalindrom(t *testing.T) {
	examples := []struct {
		num    string
		expect bool
	}{
		{num: "1", expect: true},
		{num: "2", expect: true},
		{num: "12321", expect: true},
		{num: "12322", expect: false},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %s", tt.num), func(t *testing.T) {
			got := IsPalindrom(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestReversed(t *testing.T) {
	examples := []struct {
		num    string
		expect string
	}{
		{num: "1", expect: "1"},
		{num: "2", expect: "2"},
		{num: "12321", expect: "12321"},
		{num: "12322", expect: "22321"},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %s", tt.num), func(t *testing.T) {
			got := Reversed(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	examples := []struct {
		num    int
		expect string
	}{
		{num: 1, expect: "1"},
		{num: 2, expect: "2"},
		{num: 12321, expect: "12321"},
		{num: 12322, expect: "12322"},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := ToString(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestIsLychrel(t *testing.T) {
	examples := []struct {
		num    int
		expect bool
	}{
		{num: 1, expect: false},
		{num: 2, expect: false},
		{num: 47, expect: false},
		{num: 349, expect: false},
		{num: 196, expect: true},
		{num: 4994, expect: true},
		{num: 10677, expect: true},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := IsLychrel(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestAddEm(t *testing.T) {
	examples := []struct {
		num, rev string
		expect   string
	}{
		{num: "11222", rev: "22211", expect: "33433"},
		{num: "88999", rev: "99988", expect: "188987"},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %s", tt.num), func(t *testing.T) {
			got := AddEm(tt.num, tt.rev)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}
