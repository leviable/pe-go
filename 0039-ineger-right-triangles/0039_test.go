package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestIsRightTriangle(t *testing.T) {
	examples := []struct {
		tri    Triangle
		expect bool
	}{
		{tri: Triangle{1, 1, 52}, expect: false},
		{tri: Triangle{20, 48, 52}, expect: true},
		{tri: Triangle{24, 45, 51}, expect: true},
		{tri: Triangle{30, 40, 50}, expect: true},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %v", tt.tri), func(t *testing.T) {
			got := IsRightTriangle(tt.tri)
			want := tt.expect

			if got != want {
				t.Errorf("Got %t, wanted %t", got, want)
			}
		})
	}
}

func TestSolutionGenerator(t *testing.T) {

	sol1 := Triangle{1, 1, 2}
	sol2 := Triangle{1, 1, 3}
	sol3 := Triangle{1, 1, 4}
	sol4 := Triangle{1, 2, 3}
	sol5 := Triangle{1, 1, 5}
	sol6 := Triangle{1, 2, 4}
	sol7 := Triangle{2, 2, 3}

	examples := []struct {
		perim  int
		expect []Triangle
	}{
		{perim: 4, expect: []Triangle{sol1}},
		{perim: 5, expect: []Triangle{sol2}},
		{perim: 6, expect: []Triangle{sol3, sol4}},
		{perim: 7, expect: []Triangle{sol5, sol6, sol7}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.perim), func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			got := make([]Triangle, 0)
			solGen := solutionGenerator(ctx, tt.perim)
			for sol := range solGen {
				got = append(got, sol)
			}
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestGetPerimeterSolutions(t *testing.T) {
	sol1 := Triangle{20, 48, 52}
	sol2 := Triangle{24, 45, 51}
	sol3 := Triangle{30, 40, 50}

	examples := []struct {
		perim  int
		expect []Triangle
	}{
		{perim: 3, expect: []Triangle{}},
		{perim: 120, expect: []Triangle{sol1, sol2, sol3}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.perim), func(t *testing.T) {
			got := GetPerimeterSolutions(tt.perim)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestSolutionGenerator2(t *testing.T) {

	sol1 := Triangle{3, 4, 5}
	sol2 := Triangle{6, 8, 10}
	sol3 := Triangle{5, 12, 13}

	examples := []struct {
		perim  int
		expect []Triangle
	}{
		{perim: 10, expect: []Triangle{}},
		{perim: 20, expect: []Triangle{sol1}},
		{perim: 30, expect: []Triangle{sol1, sol2, sol3}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.perim), func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			got := make([]Triangle, 0)
			solGen := solutionGenerator2(ctx, tt.perim)
			for sol := range solGen {
				got = append(got, sol)
			}
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}
