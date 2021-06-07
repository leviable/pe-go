package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestPentNumGen(t *testing.T) {
	examples := []struct {
		num    int
		expect []int
	}{
		{num: 1, expect: []int{1}},
		{num: 10, expect: []int{1, 5, 12, 22, 35, 51, 70, 92, 117, 145}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			pentNumMan := NewPentNumMan()

			got := make([]int, 0)
		loop:
			for len(got) < tt.num {
				select {
				case p := <-pentNumMan.PentNumGen(ctx):
					got = append(got, p)
				case <-ctx.Done():
					break loop
				}
			}
			cancel()

			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestQualify(t *testing.T) {
	examples := []struct {
		num1, num2 int
		expect     bool
	}{
		{num1: 22, num2: 70, expect: false},
		{num1: 1560090, num2: 7042750, expect: true},
	}

	pentNumMan := NewPentNumMan()

	for x := 0; x < 10000; x++ {
		<-pentNumMan.PentNumGen(context.Background())
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d and %d", tt.num1, tt.num2), func(t *testing.T) {

			got := pentNumMan.Qualify(tt.num1, tt.num2)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %t, wanted %t", got, want)
			}
		})
	}
}

func TestIsPentagonalNumber(t *testing.T) {
	examples := []struct {
		num    int
		expect bool
	}{
		{num: 21, expect: false},
		{num: 22, expect: true},
	}

	pentNumMan := NewPentNumMan()

	for x := 0; x < 100; x++ {
		<-pentNumMan.PentNumGen(context.Background())
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {

			got := pentNumMan.IsPentagonalNumber(tt.num)
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %t, wanted %t", got, want)
			}
		})
	}
}
