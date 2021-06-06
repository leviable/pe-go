package main

import (
	"fmt"
	"testing"
)

func TestWordValue(t *testing.T) {
	examples := []struct {
		word   string
		expect int
	}{
		{word: "sky", expect: 55},
		{word: "SKY", expect: 55},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %s", tt.word), func(t *testing.T) {
			got := WordValue(tt.word)
			want := tt.expect

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestIsTriangle(t *testing.T) {
	examples := []struct {
		num    int
		expect bool
	}{
		{num: 55, expect: true},
		{num: 54, expect: false},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := IsTriangle(tt.num)
			want := tt.expect

			if got != want {
				t.Errorf("Got %t, wanted %t", got, want)
			}
		})
	}
}
