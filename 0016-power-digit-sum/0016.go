package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func SumIntDigits(num *big.Int) (sum int) {
	numStr := num.String()
	numlen := len(numStr)
	for i := 0; i < numlen; i++ {
		dignum, _ := strconv.Atoi(string(numStr[i]))
		sum += dignum
	}
	return sum
}

func TwoPower(power int) *big.Int {
	var pow big.Int
	pow.Exp(big.NewInt(2), big.NewInt(int64(power)), nil)
	return &pow
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	num := TwoPower(1000)
	fmt.Println("Sum is ", SumIntDigits(num))
}
