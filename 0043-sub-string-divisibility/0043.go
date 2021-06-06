package main

import (
	"fmt"
	"strconv"
	"time"
)

// The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the digits 0 to 9 in some order, but it also has a rather interesting sub-string divisibility property.
// Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:
// d2d3d4=406 is divisible by 2
// d3d4d5=063 is divisible by 3
// d4d5d6=635 is divisible by 5
// d5d6d7=357 is divisible by 7
// d6d7d8=572 is divisible by 11
// d7d8d9=728 is divisible by 13
// d8d9d10=289 is divisible by 17
// Find the sum of all 0 to 9 pandigital numbers with this property.

func DivisibleBy(num string, by int) bool {
	n, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}

	return n%by == 0
}

func shift(s string) string {
	return s[1:] + string(s[0])
}

func PanDigGen(num string) <-chan string {
	c := make(chan string, 0)
	go func() {
		defer close(c)
		if len(num) <= 1 {
			c <- num
			return
		}
		for x := 0; x < len(num); x++ {
			num = shift(num)
			letter := string(num[0])
			newNum := num[1:]
			for newLetter := range PanDigGen(newNum) {
				c <- letter + newLetter
			}
		}
		return
	}()
	return c
}

func SubStringDivisible(n string) bool {
	if !DivisibleBy(n[1:4], 2) {
		return false
	} else if !DivisibleBy(n[2:5], 3) {
		return false
	} else if !DivisibleBy(n[3:6], 5) {
		return false
	} else if !DivisibleBy(n[4:7], 7) {
		return false
	} else if !DivisibleBy(n[5:8], 11) {
		return false
	} else if !DivisibleBy(n[6:9], 13) {
		return false
	} else if !DivisibleBy(n[7:10], 17) {
		return false
	}

	return true
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	sum := 0

	for p := range PanDigGen("0123456789") {
		if SubStringDivisible(p) {
			fmt.Println("Found one: ", p)
			n, _ := strconv.Atoi(p)
			sum += n
		}
	}
	fmt.Println("Sum is: ", sum)
}
