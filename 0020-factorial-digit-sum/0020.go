package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func Factorial(num int) *big.Int {
	var fact big.Int
	return fact.MulRange(int64(1), int64(num))
}

func SumDigits(num int) (sum int) {
	for _, c := range strconv.Itoa(num) {
		if cc, err := strconv.Atoi(string(c)); err == nil {
			sum += cc
		}
	}

	return sum
}

func FindFactorialSum(num int) (sum int) {
	fact := Factorial(num)

	for _, c := range fact.String() {
		if cc, err := strconv.Atoi(string(c)); err == nil {
			sum += cc
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
	fmt.Println("Sum of 100 is", FindFactorialSum(100))
}
