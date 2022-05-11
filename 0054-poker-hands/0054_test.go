package main

import (
	"reflect"
	"testing"
)

func TestLoadHand(t *testing.T) {
	got := LoadHand("5H 5C 6S 7S KD")
	want := Hand{
		Cards: []Card{
			Card{Five, Heart},
			Card{Five, Club},
			Card{Six, Spade},
			Card{Seven, Spade},
			Card{King, Diamond},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestRank(t *testing.T) {
	examples := []struct {
		name string
		hand string
		want HandRank
	}{
		{name: "RoyalFlush", hand: "TH JH QH KH AH", want: RoyalFlush},
		{name: "StraightFlush", hand: "4H 5H 6H 7H 8H", want: StraightFlush},
		{name: "FourOfAKind", hand: "5H 5D 5C AH 5S", want: FourOfAKind},
	}

	for _, tt := range examples {
		t.Run(tt.name, func(t *testing.T) {
			hand := NewHand(tt.hand)
			got := hand.Rank
			want := tt.want

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %+v, want %+v", got, want)
			}
		})
	}
}

func TestHandParser(t *testing.T) {
	examples := []struct {
		name     string
		function func(*Hand) bool
		hand     string
		want     bool
	}{
		{name: "SameSuit False", function: sameSuit, hand: "2H 4D 6C JS AH", want: false},
		{name: "SameSuit True", function: sameSuit, hand: "5H 6H 7H 8H 9H", want: true},
		{name: "TenToAce False", function: tenToAce, hand: "2H 4D 6C JS AH", want: false},
		{name: "TenToAce True", function: tenToAce, hand: "TH JH QH KH AH", want: true},
		{name: "Consecutive False", function: consecutive, hand: "2H 4D 6C JS AH", want: false},
		{name: "Consecutive True", function: consecutive, hand: "5H 6H 7H 8H 9H", want: true},
		{name: "isFourOfAKind False", function: isFourOfAKind, hand: "2H 4D 6C JS AH", want: false},
		{name: "isFourOfAKind True", function: isFourOfAKind, hand: "5H 5D 5C AH 5S", want: true},
		{name: "isFullHouse False", function: isFullHouse, hand: "2H 4D 6C JS AH", want: false},
		{name: "isFullHouse True", function: isFullHouse, hand: "5H 5D 5C AH AS", want: true},
		{name: "isFlush False", function: sameSuit, hand: "2H 4D 6C JS AH", want: false},
		{name: "isFlush True", function: sameSuit, hand: "2D 4D 6D 8D 1D", want: true},
		{name: "isStraight False", function: consecutive, hand: "2H 4D 6C JS AH", want: false},
		{name: "isStraight True", function: consecutive, hand: "2S 3C 4D 5D 6D", want: true},
		{name: "isThreeOfAKind False", function: isThreeOfAKind, hand: "2H 4D 6C JS AH", want: false},
		{name: "isThreeOfAKind True", function: isThreeOfAKind, hand: "2S 2C 2D 5D 6D", want: true},
		{name: "isTwoPair False", function: isTwoPair, hand: "2H 2D 2C JS AH", want: false},
		{name: "isTwoPair True", function: isTwoPair, hand: "2S 2C 3D 3C 6D", want: true},
		{name: "isOnePair False", function: isOnePair, hand: "2H 2D 2C JS AH", want: false},
		{name: "isOnePair True", function: isOnePair, hand: "2S 2C 3D 5D 6D", want: true},
	}

	for _, tt := range examples {
		t.Run(tt.name, func(t *testing.T) {
			hand := LoadHand(tt.hand)
			got := tt.function(&hand)
			want := tt.want

			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}
		})
	}
}

func TestHandComparator(t *testing.T) {
	examples := []struct {
		name           string
		player1        Hand
		player2        Hand
		expectedWinner int
	}{
		{name: "Hand Comparator 1", player1: NewHand("5H 5C 6S 7S KD"), player2: NewHand("2C 3S 8S 8D TD"), expectedWinner: 2},
		{name: "Hand Comparator 2", player1: NewHand("5D 8C 9S JS AC"), player2: NewHand("2C 5C 7D 8S QH"), expectedWinner: 1},
		{name: "Hand Comparator 3", player1: NewHand("2D 9C AS AH AC"), player2: NewHand("3D 6D 7D TD QD"), expectedWinner: 2},
		{name: "Hand Comparator 4", player1: NewHand("4D 6S 9H QH QC"), player2: NewHand("3D 6D 7H QD QS"), expectedWinner: 1},
		{name: "Hand Comparator 5", player1: NewHand("2H 2D 4C 4D 4S"), player2: NewHand("3C 3D 3S 9S 9D"), expectedWinner: 1},
	}

	for _, tt := range examples {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compareHands(tt.player1, tt.player2)
			want := func() Hand {
				if tt.expectedWinner == 1 {
					return tt.player1
				} else {
					return tt.player2
				}
			}()

			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %+v, want %+v", got, want)
			}
		})
	}
}

func TestGroupSort(t *testing.T) {
	examples := []struct {
		name     string
		hand     Hand
		function func(*Hand) []Card
		want     []Card
	}{
		{
			name:     "4 of a kind",
			hand:     NewHand("2C 2D 3C 2S 2H"),
			function: getFours,
			want: []Card{
				Card{Two, Club},
				Card{Two, Diamond},
				Card{Two, Spade},
				Card{Two, Heart},
				Card{Three, Club},
			},
		},
		{
			name:     "FullHouse",
			hand:     NewHand("2C 2D 3C 2S 3H"),
			function: getFullHouse,
			want: []Card{
				Card{Two, Club},
				Card{Two, Diamond},
				Card{Two, Spade},
				Card{Three, Club},
				Card{Three, Heart},
			},
		},
		{
			name:     "ThreeOfAKind",
			hand:     NewHand("2C 2D 3C 2S 4H"),
			function: getThreeOfAKind,
			want: []Card{
				Card{Two, Club},
				Card{Two, Diamond},
				Card{Two, Spade},
				Card{Four, Heart},
				Card{Three, Club},
			},
		},
		{
			name:     "TwoPair",
			hand:     NewHand("2C 2D 3C 4S 3H"),
			function: getTwoPair,
			want: []Card{
				Card{Three, Club},
				Card{Three, Heart},
				Card{Two, Club},
				Card{Two, Diamond},
				Card{Four, Spade},
			},
		},
	}

	for _, tt := range examples {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.function(&tt.hand)
			want := tt.want

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %+v, want %+v", got, want)
			}
		})
	}
}
