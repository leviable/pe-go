package main

import (
	"fmt"
	"time"
)

// The first two consecutive numbers to have two distinct prime factors are:
//
// 14 = 2 × 7
// 15 = 3 × 5
//
// The first three consecutive numbers to have three distinct prime factors are:
//
// 644 = 2² × 7 × 23
// 645 = 3 × 5 × 43
// 646 = 2 × 17 × 19.
//
// Find the first four consecutive integers to have four distinct prime factors each. What is the first of these numbers?

var primes []int

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

func PrimeFactors(num int) (factors []int) {
	if primes == nil {
		primes = GetPrimesUnder(1_000_000)
	}

	loop:
	for {
		for _, p := range primes {
			if num == p {
				factors = append(factors, p)
				break loop
			}
			if num % p == 0 {
				num /= p
				factors = append(factors, p)
				continue loop
			}
		}
	}


	return factors
}

func UniquePrimeFactors(num int) (factors []int) {
	primeFactors := PrimeFactors(num)
	loop:
	for _, p := range primeFactors {
		for _, f := range factors {
			if p == f {
				continue loop
			}
		}
		factors = append(factors, p)
	}

	return factors
}

func FindFirstSequential(num int) (first []int) {
	for x:= 9; len(first) < num; x++ {
		if x % 5000 == 0 {
			fmt.Println(x)
		}
		if len(UniquePrimeFactors(x)) == num {
			first = append(first, x)
		} else {
			first = make([]int, 0)
		}
	}
	return first
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("First 4: ", FindFirstSequential(4))
}
