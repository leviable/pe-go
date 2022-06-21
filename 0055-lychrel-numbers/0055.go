package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Reversed(num string) string {
	rev := make([]string, len(num))
	for idx := 0; idx < len(num); idx++ {
		rev[idx] = string(num[len(num)-1-idx])
	}

	return strings.Join(rev, "")
}

func ToString(num int) string {
	return strconv.Itoa(num)
}

func IsPalindrom(numStr string) bool {
	numStrRev := Reversed(numStr)

	return numStr == numStrRev
}

func IsLychrel(num int) bool {
	numStr := ToString(num)
	var iterations int
	for iterations = 0; iterations < 50; iterations++ {

		numStr = AddEm(numStr, Reversed(numStr))

		if IsPalindrom(numStr) {
			fmt.Printf("Found one (%d) (%s) after %d iterations\n", num, numStr, iterations)
			break
		}

	}

	if iterations >= 50 {
		fmt.Printf("Found one (%d) (%s) after %d iterations\n", num, numStr, iterations)
	}
	return iterations >= 48
}

func AddEm(num, rev string) string {

	newNum := make([]string, 0)
	numLen := len(num)

	var rem, n1, n2 int
	for idx := numLen - 1; idx >= 0; idx-- {
		n1, _ = strconv.Atoi(string(num[idx]))
		n2, _ = strconv.Atoi(string(rev[idx]))
		newNum = append(newNum, strconv.Itoa((n1+n2+rem)%10))
		rem = (n1 + n2 + rem) / 10

	}
	if rem > 0 {
		newNum = append(newNum, Reversed(strconv.Itoa(rem)))
	}

	return Reversed(strings.Join(newNum, ""))
}

func pe0055() (count int) {

	for x := 0; x < 10_000; x++ {
		if IsLychrel(x) {
			count++
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
	fmt.Println("Total number of Lychrel numbers under 10,000: ", pe0055())
}
