package main

import (
	"reflect"
	"testing"
)

var testNumbers = []string{
	"135740250",
	"018260538",
	"617309629",
	"154908250",
}

func TestSplitNums(t *testing.T) {
	t.Run("Split 4 digit number strings", func(t *testing.T) {
		got1, got2 := SplitNums([]string{"1234", "5678"})
		want1, want2 := []string{"123", "567"}, []string{"4", "8"}

		if !reflect.DeepEqual(got1, want1) {
			t.Errorf("Got %s, wanted %s", got1, want1)
		}

		if !reflect.DeepEqual(got2, want2) {
			t.Errorf("Got %s, wanted %s", got2, want2)
		}
	})
	t.Run("Split 1 digit number strings", func(t *testing.T) {
		got1, got2 := SplitNums([]string{"1", "5"})
		want1, want2 := []string{"", ""}, []string{"1", "5"}

		if !reflect.DeepEqual(got1, want1) {
			t.Errorf("Got %s, wanted %s", got1, want1)
		}

		if !reflect.DeepEqual(got2, want2) {
			t.Errorf("Got %s, wanted %s", got2, want2)
		}
	})
}

func TestSumLargeNumbers(t *testing.T) {
	t.Run("Sum 4 numbers", func(t *testing.T) {
		got := SumLargeNumbers(testNumbers, 0)
		want := "926218667"

		if got != want {
			t.Errorf("Got %s, wanted %s", got, want)
		}
	})

}
