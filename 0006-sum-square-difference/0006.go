package main

import (
	"fmt"
	"math"
	"time"
)

func GetSumOfSquares(maxNum int) (sum int) {
	for i := 1; i <= maxNum; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	return
}

func GetSquareOfSum(maxNum int) (sum int) {
	for i := 1; i <= maxNum; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func GetDifference(maxNum int) (difference int) {
	sumOfSquares := GetSumOfSquares(maxNum)
	squareOfSum := GetSquareOfSum(maxNum)
	return squareOfSum - sumOfSquares
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	difference := GetDifference(100)
	fmt.Printf("Difference is: %d\n", difference)
}
