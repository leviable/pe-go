package main

import (
	"fmt"
	"time"
)

type Num struct {
	num float64
	exp int
}

func ConvertNumber(rawNum int) Num {
	num := Num{num: float64(rawNum), exp: 0}

	for num.num >= 10 {
		num.num /= 10
		num.exp += 1
	}

	return num
}

func GetFactorial(num int) []Num {
	nums := make([]Num, 0)
	for x := num; x > 0; x-- {
		nums = append(nums, ConvertNumber(x))
	}
	return nums
}

func Compute(n, r int) Num {
	numerator := Num{num: 1, exp: 0}
	denominator := Num{num: 1, exp: 0}

	for _, num := range GetFactorial(n) {
		numerator.num *= num.num
		numerator.exp += num.exp
	}

	for _, den := range GetFactorial(r) {
		denominator.num *= den.num
		denominator.exp += den.exp
	}

	for _, den := range GetFactorial(n - r) {
		denominator.num *= den.num
		denominator.exp += den.exp
	}

	val := Num{}
	val.num = numerator.num / denominator.num
	val.exp = numerator.exp - denominator.exp

	return val
}

func Normalize(num Num) Num {
	for num.num >= 10 {
		num.num /= 10
		num.exp += 1
	}
	for num.num < 1 {
		num.num *= 10
		num.exp -= 1
	}
	return num
}

func IsOver(num Num, over int) bool {
	return num.exp >= over
}

func PE0053(maxN, isOver int) (total int) {

	for n := 1; n <= maxN; n++ {
		for r := 1; r <= n-1; r++ {
			if IsOver(Normalize(Compute(n, r)), isOver) {
				total++
			}
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
	fmt.Println("0053: ", PE0053(100, 6))
}
