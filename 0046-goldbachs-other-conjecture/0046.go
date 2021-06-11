package main

import (
	"fmt"
	"time"
)

// It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.
//
// 9 = 7 + 2×12
// 15 = 7 + 2×22
// 21 = 3 + 2×32
// 25 = 7 + 2×32
// 27 = 19 + 2×22
// 33 = 31 + 2×12
//
// It turns out that the conjecture was false.
//
// What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?

func GetPrimesUnder(target int) (primes []int) {
	if target <= 1 {
		return []int{}
	}
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

func squareGen(done chan struct{}) <-chan int {
	c := make(chan int, 0)

	go func() {
		defer close(c)
		for x := 1; ; x++ {
			select {
			case <-done:
				return
			case c <- 2 * x * x:
			}
		}
	}()

	return c
}

func GoldbachSieve(primes []int) (sieve []bool) {
	// Special case
	if len(primes) == 0 {
		return []bool{}
	}

	sieveSize := primes[len(primes)-1] + 1
	sieve = make([]bool, sieveSize)
	// Update all even numbers
	for idx := range sieve {
		if idx%2 == 0 {
			sieve[idx] = true
		}
	}

	sieve[1] = true

	// Update all primes
	done := make(chan struct{})
loop:
	for _, p := range primes {
		sieve[p] = true
		if p == 2 || p == 5 {
			continue
		}
		for s := range squareGen(done) {
			num := p + s
			if num > sieveSize-1 {
				done <- struct{}{}
				continue loop
			}
			sieve[num] = true
		}
	}
	return sieve
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	primes := GetPrimesUnder(10000)
	sieve := GoldbachSieve(primes)
	for idx, val := range sieve {
		if val == false {
			fmt.Println("Found one: ", idx)
		}
	}
	//fmt.Println("Sum of first 11 truncatable primes: ", foo(11))
}
