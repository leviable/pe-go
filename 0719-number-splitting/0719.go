package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//

func isSNumber(num int) bool {
	numSq := num * num
	numStr := strings.Split(strconv.Itoa(numSq), "")
	fmt.Println("************")
	fmt.Println(numStr)
	fmt.Println("************")

	c := make(chan []int, 0)
	go func(c chan<- []int) {
		for {
			c <- []int{1, 2, 3}
		}
	}(c)

	for comb := range c {
		fmt.Println("comb is: ", comb)
		sum := 0
		for _, val := range comb {
			sum += val
		}
		if sum == num {
			return true
		}
	}
	return false
}

// func GetCombos(num string) [][]string {
// 	// First determine how many digets we have
// 	numLen := len(num)
// 	return [][]string{}
// }

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	//fmt.Println("Sum of first 11 truncatable primes: ", foo(11))
}
