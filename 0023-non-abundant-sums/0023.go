package main

import (
	"fmt"
	"math"
	"time"
)

func Contains(list []int, num int) bool {
	for _, item := range list {
		if item == num {
			return true
		}
	}
	return false
}

func GetDivisors(num int) (divisors []int) {
	divisors = append(divisors, 1)
	if num == 1 || num == 2 {
		return
	}
	max := int(math.Ceil(math.Sqrt(float64(num)))) + 1
	for i := 2; i < max; i++ {
		if num%i == 0 {
			if !Contains(divisors, i) {
				divisors = append(divisors, i)
			}
			if num/i != i && !Contains(divisors, num/i) {
				divisors = append(divisors, num/i)
			}
		}
	}
	return
}

func IsAbundant(num int) bool {
	sum := 0
	for _, num := range GetDivisors(num) {
		sum += num
	}
	return sum > num
}

func GetAbundantNumbers(under int) (abNums []int) {
	for i := 2; i < under; i++ {
		if IsAbundant(i) {
			abNums = append(abNums, i)
		}
	}
	return
}

func DoIt(under int) (sum int) {
	var m = map[int]bool{}
	abNums := GetAbundantNumbers(under)
	for i, numi := range abNums {
		for _, numj := range abNums[i:] {
			tempSum := numi + numj
			if tempSum < under {
				m[tempSum] = true
			}
		}
	}

	for i := 1; i < under; i++ {
		if _, ok := m[i]; !ok {
			sum += i
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
	fmt.Println("Abundant numbers total", DoIt(28123))
}
