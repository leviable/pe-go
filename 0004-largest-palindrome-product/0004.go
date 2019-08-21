package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func IsPalindrome(num int) bool {
	numStr := strconv.Itoa(num)

	reversedRunes := make([]rune, len(numStr))
	for i, r := range numStr {
		reversedRunes[len(numStr)-1-i] = r
	}

	return numStr == string(reversedRunes)
}

func GetLargestPalindrome(numSize int) (largest int) {
	min := math.Pow10(numSize - 1)
	max := math.Pow10(numSize) - 1

	for i := max; i >= min; i-- {
		for j := max; j >= min; j-- {
			product := int(i * j)
			if IsPalindrome(product) && product > largest {
				largest = product
			}
			if product < largest {
				break
			}

		}
	}

	return
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	largestPalindrome := GetLargestPalindrome(3)
	fmt.Printf("Largest Palindrome is: %d\n", largestPalindrome)
}
