package main

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	LEFT  = "left"
	RIGHT = "right"
)

var m = map[int]bool{1: false, 2: true, 3: true, 4: false, 5: true}
var mut sync.Mutex

// The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.
// Find the sum of the only eleven primes that are both truncatable from left to right and right to left.
// NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.

func SumTruncatablePrimes(limit int) (sum int) {
	ctx, cancel := context.WithCancel(context.Background())

	count := 0
	for p := range FindTruncatablePrimes(ctx) {
		sum += p
		count++
		if count >= limit {
			cancel()
			break
		}
	}

	return
}

func FindTruncatablePrimes(ctx context.Context) <-chan int {
	primeCh := make(chan int, 0)

	go func() {
	loop:
		for x := 9; ; x += 2 {
			for t := range Truncate(x, RIGHT) {
				if !IsPrime(t) {
					continue loop
				}
			}
			for t := range Truncate(x, LEFT) {
				if !IsPrime(t) {
					continue loop
				}
			}
			select {
			case <-ctx.Done():
				break loop
			case primeCh <- x:
			}
		}

		close(primeCh)
	}()

	return primeCh
}

func Truncate(num int, direction string) <-chan int {
	intCh := make(chan int, 0)

	go func() {
		numStr := strconv.Itoa(num)
		numLen := len(numStr)

		for idx := range strings.Split(numStr, "") {
			var newNum string
			if direction == RIGHT {
				newNum = numStr[:numLen-idx]
			} else {
				newNum = numStr[idx:]
			}
			n, _ := strconv.Atoi(newNum)
			intCh <- n
		}
		close(intCh)
	}()

	return intCh
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
	mut.Lock()
	defer mut.Unlock()
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

func GeneratePrimes(ctx context.Context) <-chan int {
	primeChan := make(chan int, 0)

	go func() {
		primeChan <- 2
		primeChan <- 3
		primeChan <- 5
	loop:
		for x := 7; ; x += 2 {
			if IsPrime(x) {
				select {
				case <-ctx.Done():
					break loop
				case primeChan <- x:
				}
			}
		}
		close(primeChan)
	}()

	return primeChan
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("Sum of first 11 truncatable primes: ", SumTruncatablePrimes(11))
}
