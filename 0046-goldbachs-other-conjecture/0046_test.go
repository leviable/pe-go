package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoldbachSieve(t *testing.T) {
	examples := []struct {
		num    int
		expect []bool
	}{
		{num: 0, expect: []bool{}},
		{num: 1, expect: []bool{}},
		{num: 20, expect: []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			primes := GetPrimesUnder(tt.num)
			got := GoldbachSieve(primes)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}
