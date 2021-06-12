package main

import (
	"fmt"
	"time"
)

// The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.
//
// Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.

func Truncate(num int) int {
	return num % 10_000_000_000
}

func Expand(num int) int {
	expanded := 1
	for x:=0; x < num; x++ {
		expanded *= num
		expanded = Truncate(expanded)
	}
	return expanded
}

func FindSeries(num int) (series int) {
	for x:=1; x <= num; x++ {
		series = Truncate(series + Expand(x))
	}

	return series
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("Last 10 digits of 1000 series: ", FindSeries(1000))
}
