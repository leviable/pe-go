package main

import (
	"fmt"
	"math"
	"time"
)

func IsPrime(num int) bool {
	max := int(math.Ceil(math.Sqrt(float64(num)))) + 1
	for i := 2; i < max; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func RowGenerator() <-chan []int {
	c := make(chan []int, 0)

	go func() {
		c <- []int{1}
		for x := 1; ; x++ {
			start := (2*x - 1) * (2*x - 1)
			c <- []int{
				start + 1*(2*x),
				start + 2*(2*x),
				start + 3*(2*x),
			}
		}
	}()

	return c
}

func PE0000() int {

	gen := RowGenerator()

	cornerCount := 1
	primesFound := 0

	<-gen
	rowNum := 1
	for row := range gen {
		rowNum++
		cornerCount += 4
		for _, corner := range row {
			if IsPrime(corner) {
				// fmt.Println("Found one:", corner)
				primesFound++
			}
		}

		// fmt.Printf("Row %d (%d) val is %f\n", rowNum, cornerCount, float64(primesFound)/float64(cornerCount))
		if float64(primesFound)/float64(cornerCount) < 0.1 {
			break
		}
	}

	return 2*rowNum - 1
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("0000: ", PE0000())
}
