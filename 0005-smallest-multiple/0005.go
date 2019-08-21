package main

import (
	"fmt"
	"time"
)

func GetSmallestMultiple(num int) int {
	testNum := num
	for {
		testNum += num
		found := true
		for i := num; i >= 2; i-- {
			if testNum%i != 0 {
				found = false
				break
			}
		}
		if found == true {
			return testNum
		}
	}
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	smallestMultiple := GetSmallestMultiple(20)
	fmt.Printf("Smallest multiple for %d digits is: %d\n", 20, smallestMultiple)
}
