package main

import (
	"fmt"
	"math"
)

func GetFactors(num int) []int {
	max := int(math.Floor(float64(num) / 2.0))
	for i := 2; i <= max; i++ {
		if num%i == 0 {
			num1Factors := GetFactors(i)
			num2Factors := GetFactors(num / i)
			return append(num1Factors, num2Factors...)
		}

	}
	return []int{num}
}

func GetLargest(nums []int) (largest int) {

	for i, num := range nums {
		if i == 0 || num > largest {
			largest = num
		}
	}
	return
}

func GetLargestPrimeFactor(num int) (largest int) {
	factors := GetFactors(num)
	largest = GetLargest(factors)

	return largest
}

func main() {
	largestPrimeFactor := GetLargestPrimeFactor(600851475143)
	fmt.Printf("Largest prime factor of %d: %d", 600851475143, largestPrimeFactor)
}
