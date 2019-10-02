package main

import (
	"fmt"
	"testing"
)

func TestAddStrings(t *testing.T) {
	evalStrings := []struct {
		strA, strB, expect string
	}{
		{strA: "1", strB: "2", expect: "3"},
		{strA: "111", strB: "222", expect: "333"},
		{strA: "11", strB: "222", expect: "233"},
		{strA: "99", strB: "222", expect: "321"},
		{strA: "99", strB: "999", expect: "1098"},
	}

	for _, tt := range evalStrings {
		t.Run(fmt.Sprintf("Evaluating %q and %q", tt.strA, tt.strB), func(t *testing.T) {
			got := AddStrings(tt.strA, tt.strB)
			want := tt.expect

			if got != want {
				t.Errorf("Got %q, wanted %q", got, want)
			}
		})
	}
}

func TestFindFirstDigLen(t *testing.T) {
	got := FindFirstDigLen(3)
	want := 12

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
