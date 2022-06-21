package main

import (
	"fmt"
	"math/big"
	"time"
)

var TWO = big.NewInt(2)

func expand(count, target int) (num, den *big.Int) {
	if count == target {
		return big.NewInt(1), big.NewInt(2)
	}

	num, den = expand(count+1, target)

	newNum := big.NewInt(1)
	newNum.Mul(den, TWO)
	newNum.Add(newNum, num)
	return den, newNum
}

func Expand(target int) (num, den *big.Int) {
	num, den = expand(1, target)
	num.Add(num, den)
	return num, den
}

func IsBigger(num, den *big.Int) bool {
	numStr := num.String()
	denStr := den.String()

	return len(numStr) > len(denStr)
}

func PE0057() (sum int) {
	var num, den *big.Int
	for x := 1; x < 1_000; x++ {
		num, den = Expand(x)
		if IsBigger(num, den) {
			// fmt.Printf("Found 1 for %d/%d\n", num, den)
			sum++
		}
	}

	return
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("0057: ", PE0057())
}
