package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
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

type Primes struct {
	primes []int
}

func (p *Primes) write(prime int) {
	p.primes = append(p.primes, prime)
}

func (p *Primes) len() int {
	return len(p.primes)
}
func (p *Primes) get(nth int) int {
	return p.primes[nth-1]
}

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Waiting() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) NotWaiting() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count--
}

func (c *Counter) CurrentCount() int {
	return c.count
}

func FindNthPrimePar(nthPrime int) int {
	num := 1
	first := true
	writeChan := make(chan bool)
	primes := &Primes{}
	counter := &Counter{}

	go func(counter *Counter) {
		for {
			fmt.Printf("Current number waiting: %d\n", counter.CurrentCount())
			time.Sleep(250 * time.Millisecond)
		}
	}(counter)

	for {
		num++
		readChan := writeChan
		writeChan = make(chan bool)

		go func(writeChan, readChan chan bool, num int, primes *Primes, counter *Counter) {
			isPrime := IsPrime(num)
			counter.Waiting()
			<-readChan
			counter.NotWaiting()
			if isPrime {
				primes.write(num)
			}
			writeChan <- true
		}(writeChan, readChan, num, primes, counter)

		if num == 86028157 {
			fmt.Println("***** FINISHED *****")
		}

		if primes.len() > nthPrime {
			return primes.get(nthPrime)
		}

		if first {
			readChan <- true
			first = false
		}

	}
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	prime := FindNthPrimePar(5000001)
	fmt.Printf("Nth prime is: %d\n", prime)
}
