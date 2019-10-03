package main

import "testing"

func TestGetDistinctTerms(t *testing.T) {
	got := GetDistinctTerms(5)
	want := 15

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
