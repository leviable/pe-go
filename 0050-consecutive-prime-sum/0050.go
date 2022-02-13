package main

import (
	"fmt"
	"time"
)

func PrimeSieve(limit int) []int {
	isPrime := make([]bool, limit)
	primes := make([]int, 0)
	primes = append(primes, 2)

	for x := 3; x < limit; x += 2 {
		if isPrime[x] == true {
			continue
		}

		primes = append(primes, x)
		for y := x; y < limit; y += x {
			isPrime[y] = true
		}
	}
	return primes
}

type pMap map[int]bool

var primeMaps map[int]pMap

func IsPrime(knownPrimes []int, num int) bool {
	if primeMaps == nil {
		primeMaps = make(map[int]pMap)
	}
	pLen := len(knownPrimes)
	if m, ok := primeMaps[pLen]; !ok {
		m = make(pMap)
		for _, p := range knownPrimes {
			m[p] = true
		}
		primeMaps[pLen] = m
	}

	_, isPrime := primeMaps[pLen][num]

	return isPrime
}

func Do0050(limit int) int {
	primes := PrimeSieve(limit)
	longest := 0
	longestPrime := 0

	var primeSum, length int
	for idx := range primes {
		primeSum, length = doSum(primes, primes[idx:])
		if length > longest {
			fmt.Println("************** New Longest/Length:", primeSum, length)
			longest = length
			longestPrime = primeSum
		}
	}

	return longestPrime
}

func doSum(full, primes []int) (longest, length int) {
	var sum int
	for idx, p := range primes {
		sum += p
		if !IsPrime(full, sum) {
			continue
		} else if sum > longest {
			longest = sum
			length = idx + 1
		}
	}
	return longest, length
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("0050: ", Do0050(1_000_000))
}
