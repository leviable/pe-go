package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestManagerGen(t *testing.T) {
	examples := []struct {
		man    *Manager
		expect []int
	}{
		{man: NewTriangle(), expect: []int{1, 3, 6, 10, 15}},
		{man: NewPentagonal(), expect: []int{1, 5, 12, 22, 35}},
		{man: NewHexagonal(), expect: []int{1, 6, 15, 28, 45}},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %s", tt.man.Type), func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			got := make([]int, 0)
			for num := range tt.man.Gen(ctx) {
				got = append(got, num)
				if len(got) >= 5 {
					cancel()
					break
				}
			}
			want := tt.expect

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Got %v, wanted %v", got, want)
			}
		})
	}
}

func TestCurrent(t *testing.T) {
	examples := []struct {
		man    *Manager
		expect int
	}{
		{man: NewTriangle(), expect: 15},
		{man: NewPentagonal(), expect: 35},
		{man: NewHexagonal(), expect: 45},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %s", tt.man.Type), func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			for x := 0; x < 5; x++ {
				<-tt.man.Gen(ctx)
			}
			cancel()

			got := tt.man.Current
			want := tt.expect

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestLowest(t *testing.T) {
	examples := []struct {
		nums   []int
		expect int
	}{
		{nums: []int{1, 2, 3}, expect: 0},
		{nums: []int{2, 1, 3}, expect: 1},
		{nums: []int{2, 3, 1}, expect: 2},
		{nums: []int{1, 1, 1}, expect: -1},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("Evaluating %v", tt.nums), func(t *testing.T) {
			got := Lowest(tt.nums[0], tt.nums[1], tt.nums[2])
			want := tt.expect

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestNextMatch(t *testing.T) {
	ctx := context.Background()
	t.Run(fmt.Sprint("Evaluating NextMatch"), func(t *testing.T) {
		matchCh := NextMatch(ctx)
		got := <-matchCh
		want := 0

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}

		got = <-matchCh
		want = 1

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}

		got = <-matchCh
		want = 40755

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})

}
