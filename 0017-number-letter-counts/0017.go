package main

import (
	"fmt"
	"strings"
	"time"
)

var m = map[int]string{
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	100:  "one hundred",
	200:  "two hundred",
	300:  "three hundred",
	400:  "four hundred",
	500:  "five hundred",
	600:  "six hundred",
	700:  "seven hundred",
	800:  "eight hundred",
	900:  "nine hundred",
	1000: "one thousand",
}

func ConvertNum(num int) (numStr string) {
	if val, ok := m[num]; ok {
		numStr = val
	} else if num < 1000 && num > 100 {
		teno := num % 100
		hundo := num - teno
		if _, ok := m[teno]; ok {
			numStr = m[hundo] + " and " + m[teno]
		} else {
			uno := teno % 10
			teno := teno - uno
			numStr = m[hundo] + " and " + m[teno] + "-" + m[uno]
		}
	} else {
		uno := num % 10
		teno := num - uno
		numStr = m[teno] + "-" + m[uno]
	}

	return
}

func CountLettersInNum(num int) (count int) {
	numStr := ConvertNum(num)
	numStr = strings.Replace(numStr, " ", "", -1)
	numStr = strings.Replace(numStr, "-", "", -1)
	return len(numStr)
}

func GetTotalLetters(max int) (count int) {
	for i := 1; i <= max; i++ {
		count += CountLettersInNum(i)
	}
	return
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("Total letter count: ", GetTotalLetters(1000))
}
