package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPanDigGen(t *testing.T) {
	examples := []struct {
		num    string
		expect []string
	}{
		{num: "01", expect: []string{"10", "01"}},
		{num: "12", expect: []string{"21", "12"}},
		{num: "012", expect: []string{"102", "120", "210", "201", "021", "012"}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %s", tt.num), func(t *testing.T) {
			got := make([]string, 0)

			for panDig := range PanDigGen(tt.num) {
				got = append(got, panDig)
			}
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestDivisibleBy(t *testing.T) {
	examples := []struct {
		num    string
		by     int
		expect bool
	}{
		{num: "405", by: 2, expect: false},
		{num: "406", by: 2, expect: true},
		{num: "063", by: 3, expect: true},
		{num: "635", by: 5, expect: true},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %s", tt.num), func(t *testing.T) {
			got := DivisibleBy(tt.num, tt.by)
			want := tt.expect

			if got != want {
				t.Errorf("Got %t, wanted %t", got, want)
			}
		})
	}
}

func TestSubStringDivisible(t *testing.T) {
	examples := []struct {
		num    string
		expect bool
	}{
		{num: "1406357289", expect: true},
		{num: "1046357298", expect: false},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %s", tt.num), func(t *testing.T) {
			got := SubStringDivisible(tt.num)
			want := tt.expect

			if got != want {
				t.Errorf("Got %t, wanted %t", got, want)
			}
		})
	}
}
