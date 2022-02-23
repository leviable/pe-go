package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

var ErrNumStartsWithZero = errors.New("Invalid number: starts with zero")

func PrimeSieve(limit int) []int {
	fmt.Println("Generating primes")
	begin := time.Now()
	defer func() { fmt.Printf("Done generating primes: %s\n", time.Since(begin)) }()
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

func Replace(primes []int, num int) []int {
	masks := make([]string, 0)
	longest := make([]int, 0)
	for _, mask := range getMasks(num, masks, "") {
		found := make([]int, 0)
		for x := 0; x < 10; x++ {
			if masked, err := applyMask(num, mask, x); IsPrime(primes, masked) {
				if err != nil {
					continue
				}
				found = append(found, masked)
			}
		}

		if len(found) > len(longest) {
			longest = found
		}
	}
	return longest
}

func getMasks(num int, masks []string, current string) []string {
	numLen := int(math.Log10(float64(num))) + 1
	if len(current) >= numLen-1 {
		current += "0"
		if current != strings.Repeat("0", len(current)) {
			masks = append(masks, current)
		}
		return masks
	}
	for _, m := range []string{"0", "1"} {
		masks = getMasks(num, masks, current+m)
	}
	return masks
}

func applyMask(num int, mask string, newNum int) (int, error) {
	masked := ""
	numStr := strconv.Itoa(num)
	for idx, n := range numStr {
		if idx == 0 && newNum == 0 && string(mask[idx]) == "1" {
			return -1, ErrNumStartsWithZero
		}
		switch mask[idx] {
		case '0':
			masked += string(n)
		default:
			masked += strconv.Itoa(newNum)
		}
	}
	foo, _ := strconv.Atoi(masked)
	return foo, nil
}

func PE0051(upTo int) []int {
	var current, longest []int
	primes := PrimeSieve(upTo)
	pLen := int(math.Log10(float64(upTo))) - 1
	remove := PrimeSieve(int(math.Pow(10, float64(pLen))))
	for _, p := range primes[len(remove):] {
		current = Replace(primes, p)
		if len(current) > len(longest) {
			longest = current
		}

		if len(longest) >= 8 {
			return longest
		}
	}

	return longest
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("0051: ", PE0051(1_000_000))
}
