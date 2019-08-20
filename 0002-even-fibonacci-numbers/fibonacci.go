package main

import "fmt"

func SumEvenValuedTerms(under int) (sum int) {
	a, b := 1, 2
	sum = 2

	for {
		newB := a + b
		a = b
		b = newB

		if b > under {
			break
		} else if b%2 == 0 {
			sum += b
		}

	}
	return
}

func main() {
	under := 4000000
	sum := SumEvenValuedTerms(under)
	fmt.Printf("Even Fibonnaci numbers under %d: %d", under, sum)
}
