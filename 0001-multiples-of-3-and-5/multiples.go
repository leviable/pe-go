//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
//Find the sum of all the multiples of 3 or 5 below 1000.

package main

import "fmt"

func FindMultiples(under int) (sum int) {
	for i := 0; i < under; i++ {
		switch {
		case i%3 == 0, i%5 == 0:
			sum += i
		}
	}
	return sum
}

func main() {
	under := 1000
	sum := FindMultiples(under)
	fmt.Printf("Multiples under %d: %d", under, sum)
}
