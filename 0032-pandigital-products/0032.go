package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

const PanStr = "123456789"

func IsPandigital(a, b, c int) bool {
	asStr := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c)

	numStr := []string{}
	for _, num := range asStr {
		numStr = append(numStr, string(num))
	}
	sort.Strings(numStr)
	result := strings.Join(numStr, "")

	return result == PanStr
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func FindSumOfProducts() (sum int) {
	var product int
	panProducts := make(map[int]bool)
	for multiplicand := 2; multiplicand < 10000; multiplicand++ {
		for multiplier := multiplicand; multiplier < 10000; multiplier++ {
			product = multiplicand * multiplier
			if IsPandigital(multiplicand, multiplier, product) {
				fmt.Println("Added one: ", product)
				panProducts[product] = true
			}
		}
	}

	for prod, _ := range panProducts {
		fmt.Println("Prod is: ", prod)
		sum += prod
	}

	return
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("Sum of products:", FindSumOfProducts())
}
