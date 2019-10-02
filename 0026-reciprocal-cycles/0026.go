package main

import (
	"errors"
	"strings"
)

const MinCycles = 5

var ErrStringNotLongEnough = errors.New("string not long enough")
var ErrNoCycleFound = errors.New("could not find a cycle in number")

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}

	return
}

func GetRecurringCycle(num string) (cycle string, err error) {
	num = strings.Split(Reverse(num), ".0")[0]
	maxCycleLen := len(num) / MinCycles
	if maxCycleLen < 1 {
		return "", ErrStringNotLongEnough
	}
	for i := 1; i <= maxCycleLen; i++ {
		for j := 2; j < MinCycles; j++ {
			if num[:i] != num[i*j:i*j+i] {
				cycle = ""
				break
			} else {
				cycle = num[:i]
			}
		}

		if cycle != "" {
			cycle = Reverse(cycle)
			break
		}
	}

	if cycle == "" {
		err = ErrNoCycleFound
	}

	return cycle, err
}
