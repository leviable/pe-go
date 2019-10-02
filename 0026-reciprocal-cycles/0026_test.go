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
			got, err := GetRecurringCycle(tt.num)
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
