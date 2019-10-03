package main

import (
	"fmt"
	"time"
)

func SumDiags(width int) (sum int) {
	sum = 1
	val := 3
	for i := 3; i <= width; i += 2 {
		for j := val; j < i*i; j += (i - 1) {
			sum += j
		}
		val = i * i
	}
	sum += width * width
	return
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("Diag sum for 1001 is", SumDiags(1001))
}
