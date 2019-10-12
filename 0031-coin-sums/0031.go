package main

import (
	"fmt"
	"time"
)

const target = 200

func FindOnePence(total int) int {
	return 1
}

func FindTwoPence(total int) (count int) {
	for i := total; i <= target; i += 2 {
		count += FindOnePence(i)
	}

	return
}

func FindFivePence(total int) (count int) {
	for i := total; i <= target; i += 5 {
		count += FindTwoPence(i)
	}

	return
}

func FindTenPence(total int) (count int) {
	for i := total; i <= target; i += 10 {
		count += FindFivePence(i)
	}

	return
}

func FindTwentyPence(total int) (count int) {
	for i := total; i <= target; i += 20 {
		count += FindTenPence(i)
	}

	return
}

func FindFiftyPence(total int) (count int) {
	for i := total; i <= target; i += 50 {
		count += FindTwentyPence(i)
	}

	return
}

func FindOnePound(total int) (count int) {
	for i := total; i <= target; i += 100 {
		count += FindFiftyPence(i)
	}

	return
}

func FindTwoPound(total int) (count int) {
	for i := total; i <= target; i += 200 {
		count += FindOnePound(i)
	}

	return
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("Total combos:", FindTwoPound(0))
}
