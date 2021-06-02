package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTemplate(t *testing.T) {
	examples := []struct {
		num    int
		expect bool
	}{
		{num: 1, expect: true},
		{num: 2, expect: false},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := foo(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}
