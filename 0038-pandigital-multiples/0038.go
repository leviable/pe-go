package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Take the number 192 and multiply it by each of 1, 2, and 3:
// 192 × 1 = 192
// 192 × 2 = 384
// 192 × 3 = 576
// By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)
// The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).
// What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?

func panSort(num int) string {
	digits := strings.Split(strconv.Itoa(num), "")
	sort.Slice(digits, func(i, j int) bool { return digits[i] < digits[j] })

	return strings.Join(digits, "")

}

func IsPandigital(num int) bool {

	return panSort(num) == "123456789"
}

func PumpItUp(num int) (int, bool) {
	newNumStr := ""
	for x := 1; len(newNumStr) < 9; x++ {
		newNumStr += strconv.Itoa(num * x)
	}

	newNum, err := strconv.Atoi(newNumStr)
	return newNum, err == nil && len(newNumStr) == 9

}

func FindLargestPandigital() (largest int) {
	for x := 1; x < 333334; x++ {
		pumped, valid := PumpItUp(x)
		if valid && IsPandigital(pumped) && pumped > largest {
			largest = pumped
			fmt.Println("New largest: ", largest)
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
	fmt.Println("Largest concatenated pandigital: ", FindLargestPandigital())
}
