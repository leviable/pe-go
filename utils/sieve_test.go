package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrimeSieve(t *testing.T) {
	examples := []struct {
		num    int
		expect []int
	}{
		{num: 10, expect: []int{2, 3, 5, 7}},
		{num: 20, expect: []int{2, 3, 5, 7, 11, 13, 17, 19}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := GetPrimesUnder(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}
