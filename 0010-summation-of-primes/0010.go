package main

import (
	"fmt"
	"time"
)

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

func SumPrimes(primes *[]int) (sum int) {
	for _, prime := range *primes {
		sum += prime
	}
	return sum
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	p := GetPrimesUnder(2000000)
	sum := SumPrimes(&p)
	fmt.Printf("Sum of primes under %d is %d\n", 2000000, sum)
}
