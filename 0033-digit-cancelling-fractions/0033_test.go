package main

import (
	"fmt"
	"testing"
)

func TestCanCancelFraction(t *testing.T) {
	examples := []struct {
		num, den int
		expect   bool
	}{
		{num: 49, den: 98, expect: true},
		{num: 11, den: 12, expect: false},
		{num: 11, den: 21, expect: false},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d/%d\n", tt.num, tt.den), func(t *testing.T) {
			got := CanCancel(tt.num, tt.den)
			want := tt.expect

			if got != want {
				t.Errorf("Got %t, wanted %t", got, want)
			}
		})
	}
}
