package main

import (
	"fmt"
	"strconv"
	"time"
)

func AddStrings(strA, strB string) (sum string) {
	for len(strB) > len(strA) {
		strA = "0" + strA
	}

	remainder := 0
	for i := len(strA) - 1; i >= 0; i-- {
		strAInt, _ := strconv.Atoi(string(strA[i]))
		strBInt, _ := strconv.Atoi(string(strB[i]))
		intSum := strAInt + strBInt + remainder
		if intSum >= 10 {
			remainder = int(intSum / 10)
		} else {
			remainder = 0
		}

		strSum := strconv.Itoa(intSum)
		sum = string(strSum[len(strSum)-1]) + sum
	}
	if remainder > 0 {
		sum = strconv.Itoa(remainder) + sum
	}
	return
}

func FindFirstDigLen(digLen int) int {
	strA := "0"
	strB := "1"
	count := 1
	for len(strB) < digLen {
		strA, strB = strB, AddStrings(strA, strB)
		count++
	}
	return count
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	digLen := FindFirstDigLen(1000)
	fmt.Println("Index of fib number of length > 1000:", digLen)
}
