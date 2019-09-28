package main

import "fmt"
import "testing"

var TestNameSlice = []string{"MARY", "PATRICIA", "LINDA"}

func TestNameValue(t *testing.T) {
	testNames := []struct {
		name     string
		expected int
	}{
		{name: "MARY", expected: 57},
		{name: "PATRICIA", expected: 77},
		{name: "LINDA", expected: 40},
	}

	for _, tt := range testNames {
		t.Run(fmt.Sprintf("Evaluating %s", tt.name), func(t *testing.T) {
			got := NameValue(tt.name)
			want := tt.expected

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestGetNameScore(t *testing.T) {
	got := GetNameScore("COLIN", 938)
	want := 49714

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestGetNameScoreTotal(t *testing.T) {
	got := GetNameScoreTotal(TestNameSlice)
	want := 385

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
