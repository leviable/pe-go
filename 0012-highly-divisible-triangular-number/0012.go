package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func TriangleNumGen(writeChan chan int) {
	tNum := 0
	for i := 1; true; i++ {
		tNum += i
		writeChan <- tNum
	}
}

func GetFactors(num int) (factors []int) {
	factors = append(factors, 1)
	if num != 1 {
		factors = append(factors, num)
	}
	sqrt := math.Sqrt(float64(num))
	if sqrt == math.Ceil(sqrt) {
		factors = append(factors, int(sqrt))
	}
	for i := 2; i < int(sqrt); i++ {
		if num%i == 0 {
			factors = append(factors, i)
			factors = append(factors, num/i)
		}
	}

	sort.Ints(factors)
	return
}

func GetFirstToNumDivisors(divNum int) int {
	genChan := make(chan int)
	go TriangleNumGen(genChan)

	for num := range genChan {
		divisors := GetFactors(num)
		if len(divisors) > divNum {
			return num
		}
	}
	return 0
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	divNum := 500
	first := GetFirstToNumDivisors(divNum)
	fmt.Printf("First triangle number with over %d divisiors is: %d\n", divNum, first)
}
