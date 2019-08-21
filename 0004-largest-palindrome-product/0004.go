package main

import (
	"strconv"
)

func IsPalindrome(num int) bool {
	numStr := strconv.Itoa(num)

	reversedRunes := make([]rune, len(numStr))
	for i, r := range numStr {
		reversedRunes[len(numStr)-1-i] = r
	}

	return numStr == string(reversedRunes)
}
