package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.
// Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
// (Please note that the palindromic number, in either base, may not include leading zeros.)

func SumDBPalindromesUnder(limit int) (sum int) {

	for _, pal := range FindDBPalindromesUnder(limit) {
		sum += pal
	}

	return
}

func FindDBPalindromesUnder(limit int) []int {
	found := make([]int, 0)
	for x := 1; x < limit; x += 2 {
		if IsDoubleBasePalindrome(x) {
			found = append(found, x)
		}
	}

	return found
}

func IsDoubleBasePalindrome(num int) bool {
	numStr := strconv.Itoa(num)
	return IsPalindrome(numStr) && IsPalindrome(ToBinary(num))
}

func IsPalindrome(numStr string) bool {
	rev := ""
	if string(numStr[0]) == "0" {
		return false
	}
	for _, dig := range strings.Split(numStr, "") {
		rev = dig + rev
	}

	return rev == numStr
}

func ToBinary(num int) (bin string) {
	for num > 0 {
		div, r := num/2, num%2
		num = div
		bin = fmt.Sprintf("%d%s", r, bin)
	}
	return
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("Sum of double-base palindromes under 1,000,000: ", SumDBPalindromesUnder(1000000))
}
