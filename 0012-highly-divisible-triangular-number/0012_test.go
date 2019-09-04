package main

import (
	"reflect"
	"testing"
)

func TestTriangleNumGen(t *testing.T) {
	t.Run("Generaate 7 trinagle numbers", func(t *testing.T) {
		genChan := make(chan int)
		go TriangleNumGen(genChan)

		got := []int{}

		for i := 0; i < 10; i++ {
			got = append(got, <-genChan)
		}
		want := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})
}

func TestGetFactors(t *testing.T) {
	t.Run("Get factors of 28", func(t *testing.T) {
		got := GetFactors(28)
		want := []int{1, 2, 4, 7, 14, 28}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})
	t.Run("Get factors of 36", func(t *testing.T) {
		got := GetFactors(36)
		want := []int{1, 2, 3, 4, 6, 9, 12, 18, 36}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	})
}

func TestGetFirstToNumDivisors(t *testing.T) {
	t.Run("First over 5", func(t *testing.T) {
		got := GetFirstToNumDivisors(5)
		want := 28

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
}

func BenchmarkTrinagleNumGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genChan := make(chan int)

		go TriangleNumGen(genChan)

		got := []int{}

		for i := 0; i < 10; i++ {
			got = append(got, <-genChan)
		}
	}
}
