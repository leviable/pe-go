package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
	"time"
)

// The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways: (i) each of the three terms are prime, and, (ii) each of the 4-digit numbers are permutations of one another.
//
// There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property, but there is one other 4-digit increasing sequence.
//
// What 12-digit number do you form by concatenating the three terms in this sequence?

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

func DoIt() {
	primes3Dig := GetPrimesUnder(1_000)
	primes4Dig := GetPrimesUnder(10_000)

	primes := primes4Dig[len(primes3Dig):]

	pMap := make(map[string][]int)

	for _, p := range primes {
		pSlice := strings.Split(strconv.Itoa(p), "")
		sort.Slice(pSlice, func(i, j int) bool { return pSlice[i] < pSlice[j]})
		// fmt.Println(pSlice)
		pStr := strings.Join(pSlice, "")

		pMap[pStr] = append(pMap[pStr], p)
	}

	for _, v := range pMap {
		if len(v) < 3 {
			continue
		}

		subs := make(map[int][]string)

		for idx, p := range v {
			for x:=idx+1; x < len(v); x++ {
				sub := v[x] - p
				subs[sub] = append(subs[sub], fmt.Sprintf("%d - %d", v[x], p))
			}
		}

		for k, v := range subs {
			if len(v) > 1 {
				fmt.Println(k, v)
			}
		}
	}
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	DoIt()
	//fmt.Println("Sum of first 11 truncatable primes: ", foo(11))
}
