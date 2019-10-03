package main

import (
	"fmt"
	"math/big"
	"time"
)

var m = map[string]bool{}

func GetDistinctTerms(upto int) (count int) {
	for a := 2; a <= upto; a++ {
		for b := 2; b <= upto; b++ {
			prod := new(big.Int).Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
			m[prod.String()] = true
		}
	}
	return len(m)
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("Distinct powers up to 100 is", GetDistinctTerms(100))
}
