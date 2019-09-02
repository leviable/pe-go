package main

import (
	"fmt"
	"time"
)

type Triplet struct {
	a, b, c int
}

func GetTriplets(target int) []Triplet {
	foundTriplets := []Triplet{}
	for z := target; z >= 3; z-- {
		for y := z - 1; y >= 2; y-- {
			for x := y - 1; x >= 1; x-- {
				if IsTriplet(x, y, z) && x+y+z == target {
					foundTriplets = append(foundTriplets, Triplet{x, y, z})
				}
			}
		}
	}

	return foundTriplets
}

func IsTriplet(a, b, c int) bool {
	if a > b || b > c {
		return false
	}
	return a*a+b*b == c*c
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	triplets := GetTriplets(1000)
	fmt.Printf("Found %v", triplets)
}
