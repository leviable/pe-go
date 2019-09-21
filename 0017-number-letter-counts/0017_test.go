package main

import (
	"fmt"
	"testing"
)

func TestConvertNum(t *testing.T) {
	evalNums := []struct {
		num    int
		numStr string
	}{
		{num: 1, numStr: "one"},
		{num: 2, numStr: "two"},
		{num: 3, numStr: "three"},
		{num: 4, numStr: "four"},
		{num: 5, numStr: "five"},
		{num: 6, numStr: "six"},
		{num: 7, numStr: "seven"},
		{num: 8, numStr: "eight"},
		{num: 9, numStr: "nine"},
		{num: 10, numStr: "ten"},
		{num: 11, numStr: "eleven"},
		{num: 12, numStr: "twelve"},
		{num: 13, numStr: "thirteen"},
		{num: 14, numStr: "fourteen"},
		{num: 15, numStr: "fifteen"},
		{num: 16, numStr: "sixteen"},
		{num: 17, numStr: "seventeen"},
		{num: 18, numStr: "eighteen"},
		{num: 19, numStr: "nineteen"},
		{num: 20, numStr: "twenty"},
		{num: 22, numStr: "twenty-two"},
		{num: 30, numStr: "thirty"},
		{num: 33, numStr: "thirty-three"},
		{num: 40, numStr: "forty"},
		{num: 44, numStr: "forty-four"},
		{num: 50, numStr: "fifty"},
		{num: 55, numStr: "fifty-five"},
		{num: 60, numStr: "sixty"},
		{num: 66, numStr: "sixty-six"},
		{num: 70, numStr: "seventy"},
		{num: 77, numStr: "seventy-seven"},
		{num: 80, numStr: "eighty"},
		{num: 88, numStr: "eighty-eight"},
		{num: 90, numStr: "ninety"},
		{num: 99, numStr: "ninety-nine"},
		{num: 100, numStr: "one hundred"},
		{num: 101, numStr: "one hundred and one"},
		{num: 110, numStr: "one hundred and ten"},
		{num: 111, numStr: "one hundred and eleven"},
		{num: 200, numStr: "two hundred"},
		{num: 202, numStr: "two hundred and two"},
		{num: 212, numStr: "two hundred and twelve"},
		{num: 220, numStr: "two hundred and twenty"},
		{num: 222, numStr: "two hundred and twenty-two"},
		{num: 300, numStr: "three hundred"},
		{num: 303, numStr: "three hundred and three"},
		{num: 313, numStr: "three hundred and thirteen"},
		{num: 330, numStr: "three hundred and thirty"},
		{num: 333, numStr: "three hundred and thirty-three"},
		{num: 400, numStr: "four hundred"},
		{num: 404, numStr: "four hundred and four"},
		{num: 414, numStr: "four hundred and fourteen"},
		{num: 440, numStr: "four hundred and forty"},
		{num: 444, numStr: "four hundred and forty-four"},
		{num: 500, numStr: "five hundred"},
		{num: 505, numStr: "five hundred and five"},
		{num: 515, numStr: "five hundred and fifteen"},
		{num: 550, numStr: "five hundred and fifty"},
		{num: 555, numStr: "five hundred and fifty-five"},
		{num: 600, numStr: "six hundred"},
		{num: 606, numStr: "six hundred and six"},
		{num: 616, numStr: "six hundred and sixteen"},
		{num: 660, numStr: "six hundred and sixty"},
		{num: 666, numStr: "six hundred and sixty-six"},
		{num: 700, numStr: "seven hundred"},
		{num: 707, numStr: "seven hundred and seven"},
		{num: 717, numStr: "seven hundred and seventeen"},
		{num: 770, numStr: "seven hundred and seventy"},
		{num: 777, numStr: "seven hundred and seventy-seven"},
		{num: 800, numStr: "eight hundred"},
		{num: 808, numStr: "eight hundred and eight"},
		{num: 818, numStr: "eight hundred and eighteen"},
		{num: 880, numStr: "eight hundred and eighty"},
		{num: 888, numStr: "eight hundred and eighty-eight"},
		{num: 900, numStr: "nine hundred"},
		{num: 909, numStr: "nine hundred and nine"},
		{num: 919, numStr: "nine hundred and nineteen"},
		{num: 990, numStr: "nine hundred and ninety"},
		{num: 999, numStr: "nine hundred and ninety-nine"},
		{num: 1000, numStr: "one thousand"},
	}

	for _, tt := range evalNums {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := ConvertNum(tt.num)
			want := tt.numStr

			if got != want {
				t.Errorf("Got %s, wanted %s", got, want)
			}
		})
	}
}

func TestCountLettersInNum(t *testing.T) {
	evalNums := []struct {
		num, numCount int
	}{
		{num: 115, numCount: 20},
		{num: 342, numCount: 23},
	}

	for _, tt := range evalNums {
		t.Run(fmt.Sprintf("Evaluating %d", tt.num), func(t *testing.T) {
			got := CountLettersInNum(tt.num)
			want := tt.numCount

			if got != want {
				t.Errorf("Got %d, wanted %d", got, want)
			}
		})
	}
}

func TestGetTotalLetters(t *testing.T) {
	got := GetTotalLetters(5)
	want := 19

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
