package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"time"
)

type CardValue int
type CardSuit int

const (
	Unknown CardValue = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace

	Club CardSuit = iota
	Spade
	Heart
	Diamond
)

var ValueMap = map[string]CardValue{
	"2": Two,
	"3": Three,
	"4": Four,
	"5": Five,
	"6": Six,
	"7": Seven,
	"8": Eight,
	"9": Nine,
	"T": Ten,
	"J": Jack,
	"Q": Queen,
	"K": King,
	"A": Ace,
}

var SuitMap = map[string]CardSuit{
	"C": Club,
	"S": Spade,
	"H": Heart,
	"D": Diamond,
}

type HandRank int

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

type Hand struct {
	Cards    []Card
	Rank     HandRank
	WinOrder []Card
}

func NewHand(handStr string) Hand {
	hand := LoadHand(handStr)
	hand.determineRank()

	return hand
}

func (h *Hand) determineRank() {
	if sameSuit(h) && tenToAce(h) {
		// Royal Flush: Ten, Jack, Queen, King, Ace, in same suit.
		h.Rank = RoyalFlush
		h.WinOrder = []Card{h.Cards[4], h.Cards[3], h.Cards[2], h.Cards[1], h.Cards[0]}
	} else if sameSuit(h) && consecutive(h) {
		// Straight Flush: All cards are consecutive values of same suit.
		h.Rank = StraightFlush
		h.WinOrder = []Card{h.Cards[4], h.Cards[3], h.Cards[2], h.Cards[1], h.Cards[0]}
	} else if isFourOfAKind(h) {
		// Four of a Kind: Four cards of the same value.
		h.Rank = FourOfAKind
		h.WinOrder = getFours(h)
	} else if isFullHouse(h) {
		// Full House: Three of a kind and a pair.
		h.Rank = FullHouse
		h.WinOrder = getFullHouse(h)
	} else if sameSuit(h) {
		// Flush: All cards of the same suit.
		h.Rank = Flush
		h.WinOrder = []Card{h.Cards[4], h.Cards[3], h.Cards[2], h.Cards[1], h.Cards[0]}
	} else if consecutive(h) {
		// Straight: All cards are consecutive values.
		h.Rank = Straight
		h.WinOrder = []Card{h.Cards[4], h.Cards[3], h.Cards[2], h.Cards[1], h.Cards[0]}
	} else if isThreeOfAKind(h) {
		// Three of a Kind: Three cards of the same value.
		h.Rank = ThreeOfAKind
		h.WinOrder = getThreeOfAKind(h)
	} else if isTwoPair(h) {
		// Two Pairs: Two different pairs.
		h.Rank = TwoPair
		h.WinOrder = getTwoPair(h)
	} else if isOnePair(h) {
		// One Pair: Two cards of the same value.
		h.Rank = OnePair
		h.WinOrder = getOnePair(h)
	} else {
		// High Card: Highest value card.
		h.Rank = HighCard
		h.WinOrder = []Card{h.Cards[4], h.Cards[3], h.Cards[2], h.Cards[1], h.Cards[0]}
	}
}

func sameSuit(hand *Hand) bool {
	suit := hand.Cards[0].Suit
	for _, c := range hand.Cards[1:] {
		if c.Suit != suit {
			return false
		}
	}

	return true
}

func tenToAce(hand *Hand) bool {
tenToAceLoop:
	for _, royalFlushCard := range []CardValue{Ten, Jack, Queen, King, Ace} {
		for _, card := range hand.Cards {
			if card.Value == royalFlushCard {
				continue tenToAceLoop
			}
		}
		return false
	}

	return true
}

func consecutive(hand *Hand) bool {
	card := hand.Cards[0]
	for _, next := range hand.Cards[1:] {
		if card.Value+1 == next.Value {
			card = next
			continue
		}
		return false
	}
	return true
}

func isFourOfAKind(hand *Hand) bool {
	cardCount := make(map[CardValue]int)
	for _, card := range hand.Cards {
		cardCount[card.Value]++
	}

	for _, v := range cardCount {
		if v == 4 {
			return true
		}
	}

	return false
}

func getFours(hand *Hand) []Card {
	cardCount := make(map[CardValue]int)
	for _, card := range hand.Cards {
		cardCount[card.Value]++
	}

	var found CardValue
	for k, v := range cardCount {
		if v == 4 {
			found = k
			break
		}
	}

	foak := make([]Card, 0)
	extras := make([]Card, 0)

	for _, card := range hand.Cards {
		if card.Value == found {
			foak = append(foak, card)
		} else {
			extras = append(extras, card)
		}
	}

	return append(foak, extras...)
}

func getFullHouse(hand *Hand) []Card {
	cardCount := make(map[CardValue]int)
	for _, card := range hand.Cards {
		cardCount[card.Value]++
	}

	var found3 CardValue
	for k, v := range cardCount {
		if v == 3 {
			found3 = k
		}
	}

	toak := make([]Card, 0)
	onePair := make([]Card, 0)

	for _, card := range hand.Cards {
		if card.Value == found3 {
			toak = append(toak, card)
		} else {
			onePair = append(onePair, card)
		}
	}

	return append(toak, onePair...)
}

func isFullHouse(hand *Hand) bool {
	return isOnePair(hand) && isThreeOfAKind(hand)
}

func getThreeOfAKind(hand *Hand) []Card {
	cardCount := make(map[CardValue]int)
	for _, card := range hand.Cards {
		cardCount[card.Value]++
	}

	var found3 CardValue
	for k, v := range cardCount {
		if v == 3 {
			found3 = k
		}
	}

	toak := make([]Card, 0)
	onePair := make([]Card, 0)

	for _, card := range hand.Cards {
		if card.Value == found3 {
			toak = append(toak, card)
		} else {
			onePair = append(onePair, card)
		}
	}

	return append(toak, onePair[1], onePair[0])
}

func isThreeOfAKind(hand *Hand) bool {
	cardCount := make(map[CardValue]int)
	for _, card := range hand.Cards {
		cardCount[card.Value]++
	}
	for _, v := range cardCount {
		if v == 3 {
			return true
		}
	}

	return false
}

func getTwoPair(hand *Hand) []Card {
	// First: Count the card types
	cardCount := make(map[CardValue]int)
	for _, card := range hand.Cards {
		cardCount[card.Value]++
	}

	// Sort the found types, looking for two pairs
	found1, found2, highCard := Unknown, Unknown, Unknown
	for k, v := range cardCount {
		if v == 2 && found1 == Unknown {
			found1 = k
		} else if v == 2 && found2 == Unknown {
			found2 = k
		} else if v == 1 && highCard == Unknown {
			highCard = k
		}
	}

	// Compare found pairs and sort for largest/smallest
	twoPair := make([]Card, 0)

	var big, small CardValue

	if found1 > found2 {
		big = found1
		small = found2
	} else {
		big = found2
		small = found1
	}

	for _, card := range hand.Cards {
		if card.Value == big {
			twoPair = append(twoPair, card)
		}
	}

	for _, card := range hand.Cards {
		if card.Value == small {
			twoPair = append(twoPair, card)
		}
	}

	for _, card := range hand.Cards {
		if card.Value == highCard {
			twoPair = append(twoPair, card)
		}
	}

	// Append larger pair, smaller pair, and then high card
	return twoPair
}

func isTwoPair(hand *Hand) bool {
	cardCount := make(map[CardValue]int)
	for _, card := range hand.Cards {
		cardCount[card.Value]++
	}

	pairCount := 0
	for _, v := range cardCount {
		if v == 2 {
			pairCount++
		}
	}

	return pairCount == 2
}

func getOnePair(hand *Hand) []Card {
	cardCount := make(map[CardValue]int)
	for _, card := range hand.Cards {
		cardCount[card.Value]++
	}

	var found1 CardValue
	for k, v := range cardCount {
		if v == 2 {
			found1 = k
		}
	}

	onePair := make([]Card, 0)
	extra := make([]Card, 0)

	for _, card := range hand.Cards {
		if card.Value == found1 {
			onePair = append(onePair, card)
		} else {
			extra = append(extra, card)
		}
	}

	return append(onePair, extra[2], extra[1], extra[0])
}

func isOnePair(hand *Hand) bool {
	cardCount := make(map[CardValue]int)
	for _, card := range hand.Cards {
		cardCount[card.Value]++
	}
	for _, v := range cardCount {
		if v == 2 {
			return true
		}
	}

	return false
}

type Card struct {
	Value CardValue
	Suit  CardSuit
}

func LoadHand(hand string) Hand {
	cards := make([]Card, 5)
	for idx, card := range strings.Split(hand, " ") {
		cards[idx] = Card{ValueMap[string(card[0])], SuitMap[string(card[1])]}
	}
	sort.SliceStable(cards, func(i, j int) bool {
		return cards[i].Value < cards[j].Value
	})

	return Hand{Cards: cards, WinOrder: make([]Card, 0)}
}

func compareHands(player1, player2 Hand) (Hand, error) {
	if player1.Rank > player2.Rank {
		return player1, nil
	} else if player1.Rank < player2.Rank {
		return player2, nil
	} else if player1.Rank == player2.Rank {
		for x := 0; x < 5; x++ {
			if player1.WinOrder[x].Value > player2.WinOrder[x].Value {
				return player1, nil
			} else if player1.WinOrder[x].Value < player2.WinOrder[x].Value {
				return player2, nil
			} else {
				continue
			}
		}
	}
	return player1, errors.New("Could not find a clear winner")
}

type Battle struct {
	Player1 Hand
	Player2 Hand
}

func LoadPlayerHands() []Battle {
	battles := make([]Battle, 0)

	lines, _ := readLines()
	for _, line := range lines {
		cards := strings.Split(line, " ")
		battles = append(battles, Battle{
			Player1: NewHand(strings.Join(cards[:5], " ")),
			Player2: NewHand(strings.Join(cards[5:], " ")),
		})
	}

	return battles
}

func readLines() ([]string, error) {
	file, err := os.Open("poker.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func PE0000() int {

	battles := LoadPlayerHands()

	player1Wins := 0
	for idx, battle := range battles {
		winner, err := compareHands(battle.Player1, battle.Player2)

		if err != nil {
			panic(err)
		}

		if reflect.DeepEqual(winner, battle.Player1) {
			player1Wins++
		} else {
		}
	}

	return player1Wins
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("0000: ", PE0000())
}
