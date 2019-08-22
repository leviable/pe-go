package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func IsCandidate(candidate int) bool {
	candidateStr := strconv.Itoa(candidate)
	lastChar := candidateStr[len(candidateStr)-1:]
	switch lastChar {
	case "1", "3", "7", "9":
		return true
	default:
		return false
	}
}

func IsPrime(candidate int) bool {
	switch candidate {
	case 1:
		return false
	case 2, 5:
		return true
	default:
		if !IsCandidate(candidate) {
			return false
		}
		max := int(math.Ceil(math.Sqrt(float64(candidate)))) + 1
		for i := 2; i < max; i++ {
			if candidate%i == 0 {
				return false
			}
		}
	}
	return true
}

func FindNthPrime(nthPrime int) int {
	start := 1
	primeCount := 0
	for {
		start++
		if IsPrime(start) {
			primeCount++
		}

		if primeCount == nthPrime {
			return start
		}

	}
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	prime := FindNthPrime(10001)
	fmt.Printf("Nth prime is: %d\n", prime)
}
