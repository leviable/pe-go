package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const MinCycles = 5

var ErrStringNotLongEnough = errors.New("string not long enough")
var ErrNoCycleFound = errors.New("could not find a cycle in number")
var ErrNotRepeatingDecimal = errors.New("not a repeating decimal")

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}

	return
}

func DetectCycle(num string) (cycle string, err error) {
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

func FindCycleForDenominator(denom int) (cycle string, err error) {
	num := "0."
	remainder := 1
	for remainder > 0 {
		val := strconv.Itoa((remainder * 10) / denom)
		num = num + val
		remainder = (remainder * 10) % denom

		cycle, err = DetectCycle(num)
		if err != nil {
			continue
		} else {
			break
		}
	}

	if remainder == 0 {
		err = ErrNotRepeatingDecimal
	}

	return
}

func FindLongestRecurringCycle(upto int) (longestDenom int) {
	longest := 0
	for i := 3; i < upto; i += 2 {
		cycle, err := FindCycleForDenominator(i)
		if err != nil {
			continue
		}
		cycleLen := len(cycle)
		if cycleLen > longest {
			longest = cycleLen
			longestDenom = i
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
	longest := FindLongestRecurringCycle(1000)
	fmt.Println("Denom with longest recurring cycle is", longest)
}
