package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// An irrational decimal fraction is created by concatenating the positive integers:
//		0.123456789101112131415161718192021...
// It can be seen that the 12th digit of the fractional part is 1.
// If dn represents the nth digit of the fractional part, find the value of the following expression.
// d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000

func DigitGenerator(limit int) <-chan int {
	digCh := make(chan int, 0)

	go func() {
		defer close(digCh)
		count := 0
		for x := 1; count < limit; x++ {
			xStr := strings.Split(strconv.Itoa(x), "")
			for i := 0; count < limit && i < len(xStr); i++ {
				dInt, _ := strconv.Atoi(xStr[i])
				digCh <- dInt
				count++
			}
		}
	}()

	return digCh
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	val := 1

	gen := DigitGenerator(1000001)

	for x := 1; x < 1000001; x++ {
		dig := <-gen
		for _, d := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
			if x == d {
				val *= dig
			}
		}
	}
	fmt.Println("Solution is: ", val)
	//fmt.Println("Sum of first 11 truncatable primes: ", foo(11))
}
