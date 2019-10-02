package main

import (
	"fmt"
	"math"
	"time"
)

var m = map[int]bool{1: false, 2: true, 3: true, 4: false, 5: true}

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

func IsCandidate(candidate int) bool {
	switch candidate % 10 {
	case 1, 3, 7, 9:
		return true
	default:
		return false
	}
}

func IsPrime(candidate int) bool {
	if val, ok := m[candidate]; ok {
		return val
	}

	if !IsCandidate(candidate) {
		m[candidate] = false
		return false
	}
	max := int(math.Ceil(math.Sqrt(float64(candidate)))) + 1
	for i := 2; i < max; i++ {
		if candidate%i == 0 {
			m[candidate] = false
			return false
		}
	}

	m[candidate] = true
	return true
}

func GetPrimeList(a, b int) (primes []int) {
	for i := 0; true; i++ {
		candidate := i*i + a*i + b
		if !IsPrime(candidate) {
			break
		}

		primes = append(primes, candidate)
	}
	return
}

func FindMostPrimes() (coefProd int) {
	// Pre-populate b primes
	bPrimes := GetPrimesUnder(1000)

	for _, b := range bPrimes {
		m[b] = true
	}

	most := 0
	for a := -1000; a <= 1000; a++ {
		for _, b := range bPrimes {
			coef := a * b
			primeCount := len(GetPrimeList(a, b))
			if primeCount > most {
				most = primeCount
				coefProd = coef
			}
		}
	}

	return coefProd
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	most := FindMostPrimes()
	fmt.Println("Product of coefficients with most primes is", most)
}
