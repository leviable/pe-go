package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenPermutations(t *testing.T) {
	got := GenPermutations([]int{0, 1, 2})
	want := []string{"012", "021", "102", "120", "201", "210"}

	sort.Strings(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
