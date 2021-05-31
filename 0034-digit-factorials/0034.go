package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.
// Find the sum of all numbers which are equal to the sum of the factorial of their digits.
// Note: As 1! = 1 and 2! = 2 are not sums they are not included.

func IsDigitFactorial(num int) bool {
	sum := 0

	for _, dig := range SplitNum(num) {
		sum += Factorial(dig)
	}
	return sum == num
}

func SplitNum(num int) []int {
	split := make([]int, 0)

	for _, dig := range strings.Split(strconv.Itoa(num), "") {
		d, _ := strconv.Atoi(dig)
		split = append(split, d)
	}

	return split
}

func Factorial(num int) (fac int) {
	if num == 0 {
		return 1
	} else if num == 1 || num == 2 {
		return num
	}

	fac = 2
	for x := 3; x < num+1; x++ {
		fac *= x
	}

	return
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	num := 9
	for {
		num++
		if IsDigitFactorial(num) {
			fmt.Println("Found one: ", num)
		}
	}
	//fmt.Println("Fifth power sum is", SumAllForPower(5))
}
