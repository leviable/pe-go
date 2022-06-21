package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Mult(num, pow int) []int {
	numStr := strings.Split(strconv.Itoa(num), "")
	nums := make([]int, len(numStr))

	for x := 0; x < len(numStr); x++ {
		nums[x], _ = strconv.Atoi(numStr[x])
	}

	var n1, n2, rem int
	for x := 1; x < pow; x++ {
		rem = 0
		newNums := make([]int, 0)
		for y := len(nums) - 1; y >= 0; y-- {
			n1 = nums[y]
			n2 = ((n1 * num) + rem) % 10
			rem = ((n1 * num) + rem) / 10

			newNums = append([]int{n2}, newNums...)
		}

		if rem > 0 {
			remStr := strings.Split(strconv.Itoa(rem), "")
			for idx := len(remStr) - 1; idx >= 0; idx-- {

				newNum, _ := strconv.Atoi(remStr[idx])
				newNums = append([]int{newNum}, newNums...)
			}
		}

		nums = newNums
	}
	return nums
}

func SumArray(num []int) (sum int) {
	for _, x := range num {
		sum += x
	}
	return
}

func PE0056() int {

	max := 0
	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			newSum := SumArray(Mult(a, b))

			if newSum > max {
				max = newSum
			}
		}
	}

	return max
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("0056: ", PE0056())
}
