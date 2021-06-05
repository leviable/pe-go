package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestDigitGenerator(t *testing.T) {
	examples := []struct {
		limit  int
		expect string
	}{
		{limit: 1, expect: "1"},
		{limit: 2, expect: "12"},
		{limit: 12, expect: "123456789101"},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.limit), func(t *testing.T) {
			got := ""

			for dig := range DigitGenerator(tt.limit) {
				got += strconv.Itoa(dig)
			}
			want := tt.expect

			if got != want {
				t.Errorf("Got %s, wanted %s", got, want)
			}
		})
	}
}
