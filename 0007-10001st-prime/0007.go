package main

import (
	"math"
	"strconv"
)

func IsCandidate(candidate int) bool {
	candidateStr := strconv.Itoa(candidate)
	lastChar := candidateStr[len(candidateStr)-1:]
	switch lastChar {
	case "1", "3", "7", "9":
		return true
	default:
		return false
	}
}

func IsPrime(candidate int) bool {
	switch candidate {
	case 1:
		return false
	case 2, 5:
		return true
	default:
		if !IsCandidate(candidate) {
			return false
		}
		max := int(math.Ceil(math.Sqrt(float64(candidate)))) + 1
		for i := 2; i < max; i++ {
			if candidate%i == 0 {
				return false
			}
		}
	}
	return true
}
