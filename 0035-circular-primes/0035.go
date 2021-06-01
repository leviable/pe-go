package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.
// There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
// How many circular primes are there below one million?

func IsCircularPrime(prime int, primes []int) bool {
loop:
	for p := range Rotate(prime) {
		for _, x := range primes {
			if x == p {
				continue loop
			}
		}
		return false
	}
	return true
}

func PrunePrimes(primes []int) []int {
	pruned := make([]int, 0)
pLoop:
	for _, p := range primes {
		if p == 2 || p == 3 || p == 5 {
			pruned = append(pruned, p)
			continue
		}
		for _, dig := range strings.Split(strconv.Itoa(p), "") {
			if dig == "1" || dig == "3" || dig == "7" || dig == "9" {
				continue
			}
			continue pLoop
		}
		pruned = append(pruned, p)
	}

	return pruned
}

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

func Rotate(num int) <-chan int {
	rotated := make(chan int)
	numStr := strconv.Itoa(num)
	go func() {
		for x := 0; x < len(numStr); x++ {
			numStr = numStr[1:] + string(numStr[0])
			n, _ := strconv.Atoi(numStr)
			rotated <- n
		}
		close(rotated)
	}()

	return rotated
}

func GetCircularPrimesUnder(limit int) []int {
	circularPrimes := make([]int, 0)

	largerSieve := PrunePrimes(GetPrimesUnder(limit * 10))
	pruned := PrunePrimes(GetPrimesUnder(limit))
	for _, prime := range pruned {
		if IsCircularPrime(prime, largerSieve) {
			circularPrimes = append(circularPrimes, prime)
		}
	}

	return circularPrimes
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("Circular Primes under 1,000,000: ", len(GetCircularPrimesUnder(1000000)))
}
