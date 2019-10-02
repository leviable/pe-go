package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	got := Reverse("123456")
	want := "654321"

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}

func TestGetRecurringCycle(t *testing.T) {

	evalStrings := []struct {
		num, expect string
		err         error
	}{
		{num: "0.1666", expect: "", err: ErrStringNotLongEnough},
		{num: "0.16666", expect: "", err: ErrNoCycleFound},
		{num: "0.166666", expect: "6", err: nil},
		{num: "0.123123123123", expect: "", err: ErrNoCycleFound},
		{num: "0.123123123123123", expect: "123", err: nil},
		{num: "0.1123123123123123", expect: "123", err: nil},
	}

	for _, tt := range evalStrings {
		t.Run(fmt.Sprintf("Evaluating %q", tt.num), func(t *testing.T) {
			got, err := DetectCycle(tt.num)
			want := tt.expect

			if err != tt.err {
				t.Fatal(fmt.Sprintf("Got error %v, wanted error %v", err, tt.err))
			}

			if got != want {
				t.Errorf("Got %q, wanted %q", got, want)
			}
		})
	}
}

func TestFindCycleForDenominator(t *testing.T) {
	denominators := []struct {
		denom  int
		expect string
		err    error
	}{
		{denom: 2, expect: "", err: ErrNotRepeatingDecimal},
		{denom: 3, expect: "3", err: nil},
		{denom: 4, expect: "", err: ErrNotRepeatingDecimal},
		{denom: 5, expect: "", err: ErrNotRepeatingDecimal},
		{denom: 6, expect: "6", err: nil},
		{denom: 7, expect: "142857", err: nil},
		{denom: 8, expect: "", err: ErrNotRepeatingDecimal},
		{denom: 9, expect: "1", err: nil},
	}

	for _, tt := range denominators {
		t.Run(fmt.Sprintf("Evaluating %d", tt.denom), func(t *testing.T) {
			got, err := FindCycleForDenominator(tt.denom)
			want := tt.expect

			if err != tt.err {
				t.Fatal(fmt.Sprintf("Got error %v, wanted error %v", err, tt.err))
			}

			if got != want {
				t.Errorf("Got %q, wanted %q", got, want)
			}
		})
	}
}

func TestFindLongestRecurringCycle(t *testing.T) {
	got := FindLongestRecurringCycle(10)
	want := 7

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
