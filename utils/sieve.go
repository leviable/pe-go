package main

import "fmt"

func main() {
	fmt.Println("vim-go")
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
