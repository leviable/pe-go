package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Pentagonal numbers are generated by the formula, Pn=n(3n−1)/2. The first ten pentagonal numbers are:
//				1, 5, 12, 22, 35, 51, 70, 92, 117, 145, ...
// It can be seen that P4 + P7 = 22 + 70 = 92 = P8. However, their difference, 70 − 22 = 48, is not pentagonal.
// Find the pair of pentagonal numbers, Pj and Pk, for which their sum and difference are pentagonal and D = |Pk − Pj| is minimised; what is the value of D?

type PentNumMan struct {
	pentNums map[int]bool
	genCh    chan int
	mu       sync.Mutex
}

func NewPentNumMan() *PentNumMan {
	return &PentNumMan{
		pentNums: make(map[int]bool),
	}
}

func (p *PentNumMan) PentNumGen(ctx context.Context) <-chan int {
	if p.genCh == nil {
		p.genCh = make(chan int, 0)

		go func() {
			defer close(p.genCh)
			count := 1
		loop:
			for {
				pn := (count * (3*count - 1)) / 2
				select {
				case <-ctx.Done():
					break loop
				case p.genCh <- pn:
					p.mu.Lock()
					p.pentNums[pn] = true
					p.mu.Unlock()
					count++
				}
			}
		}()
	}

	return p.genCh
}

func (p *PentNumMan) Qualify(num1, num2 int) bool {
	a := num1 + num2
	b := num2 - num1
	return p.IsPentagonalNumber(a) && p.IsPentagonalNumber(b)
}

func (p *PentNumMan) IsPentagonalNumber(num int) bool {
	p.mu.Lock()
	_, ok := p.pentNums[num]
	p.mu.Unlock()
	return ok
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	p := NewPentNumMan()
	nums := make([]int, 0)
	for x := 0; x < 100000; x++ {
		nums = append(nums, <-p.PentNumGen(context.Background()))
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			n1 := nums[i]
			n2 := nums[j]
			if p.Qualify(n1, n2) {
				fmt.Printf("%d and %d\n", n1, n2)
			}
		}
	}
	//fmt.Println("Sum of first 11 truncatable primes: ", foo(11))
}
