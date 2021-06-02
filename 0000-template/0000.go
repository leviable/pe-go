package main

import (
	"fmt"
	"time"
)

//

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("Sum of first 11 truncatable primes: ", SumTruncatablePrimes(11))
}
