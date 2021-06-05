package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once.
// For example, 2143 is a 4-digit pandigital and is also prime.
// What is the largest n-digit pandigital prime that exists?

func GetPrimesUnder(target int) (primes []int) {
	notPrimes := make([]bool, target+1)
	notPrimes[0] = true
	notPrimes[1] = true
	notPrimes[2] = false
	primes = append(primes, 2)
	for i := 3; i < target+1; i += 2 {
		if notPrimes[i] == false {
			primes = append(primes, i)
			for j := i + i; j < target+1; j += i {
				if j > target {
					break
				}
				notPrimes[j] = true
			}
		}
	}
	return primes
}

func panSort(num int) string {
	digits := strings.Split(strconv.Itoa(num), "")
	sort.Slice(digits, func(i, j int) bool { return digits[i] < digits[j] })

	return strings.Join(digits, "")

}

func IsPandigital(num int) bool {
	numStr := strconv.Itoa(num)

	return panSort(num) == "123456789"[:len(numStr)]
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("Generating primes")
	primes := GetPrimesUnder(987654322)
	fmt.Println("Primes generated: ", len(primes))
loop:
	for x := len(primes) - 1; x > 0; x-- {
		fmt.Println("Prime is ", primes[x])
		if IsPandigital(primes[x]) {
			fmt.Println("Largest: ", primes[x])
			break loop
		}
	}
	//fmt.Println("Sum of first 11 truncatable primes: ", foo(11))
}
