package main

import "fmt"
import "reflect"
import "testing"

func TestIsTriplet(t *testing.T) {
	potentialTriplets := []struct {
		a, b, c int
		expect  bool
	}{
		{a: 3, b: 4, c: 5, expect: true},
		{a: 5, b: 12, c: 13, expect: true},
		{a: 4, b: 3, c: 5, expect: false},
		{a: 3, b: 5, c: 4, expect: false},
	}

	for _, tt := range potentialTriplets {
		t.Run(fmt.Sprintf("Evaluating %d, %d, %d", tt.a, tt.b, tt.c), func(t *testing.T) {
			got := IsTriplet(tt.a, tt.b, tt.c)
			want := tt.expect
			if got != want {
				t.Errorf("For %d-%d-%d got %t, want %t", tt.a, tt.b, tt.b, got, want)
			}
		})
	}
}

func TestFindTriplets(t *testing.T) {
	t.Run("Finding triplets equal to 12", func(t *testing.T) {
		got := GetTriplets(12)
		want := []Triplet{Triplet{3, 4, 5}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
