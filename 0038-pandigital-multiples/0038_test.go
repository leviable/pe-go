package main

import (
	"fmt"
	"testing"
)

func TestPumpItUp(t *testing.T) {
	examples := []struct {
		num    int
		expect int
		valid  bool
	}{
		{num: 1, expect: 123456789, valid: true},
		{num: 2, expect: 2468101214, valid: false},
		{num: 192, expect: 192384576, valid: true},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got, valid := PumpItUp(tt.num)
			if tt.valid != valid {
				t.Fatalf("Got invalid num back: %d", got)
			}
			want := tt.expect

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestIsPandigital(t *testing.T) {
	examples := []struct {
		num    int
		expect bool
	}{
		{num: 123456789, expect: true},
		{num: 2468101214, expect: false},
		{num: 192384576, expect: true},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := IsPandigital(tt.num)
			want := tt.expect

			if got != want {
				t.Errorf("Got %t, wanted %t", got, want)
			}
		})
	}
}
