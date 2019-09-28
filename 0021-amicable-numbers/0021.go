package main

import (
	"fmt"
	"math"
	"time"
)

var m = map[int]int{}

func GetDivisors(num int) (divisors []int) {
	divisors = append(divisors, 1)
	if num == 1 || num == 2 {
		return
	}
	max := int(math.Ceil(math.Sqrt(float64(num)))) + 1
	for i := 2; i < max; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
			if num/i != i {
				divisors = append(divisors, num/i)
			}
		}
	}
	return
}

func GetSum(num int) (sum int) {
	if val, ok := m[num]; ok {
		sum = val
	} else {
		for _, div := range GetDivisors(num) {
			sum += div
		}
		m[num] = sum
	}

	return
}

func HasAmicable(num int) bool {
	sum := GetSum(num)
	amicableSum := GetSum(sum)

	return sum != amicableSum && num == amicableSum
}

func FindSumAmicableNums(under int) (sum int) {
	amicable := []int{}

	for i := 3; i < under; i++ {
		if HasAmicable(i) {
			amicable = append(amicable, i)
		}
	}

	for _, num := range amicable {
		sum += num
	}
	return
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	sum := FindSumAmicableNums(10000)
	fmt.Println("Sum of amicable numbers under 10,000 is", sum)
}
