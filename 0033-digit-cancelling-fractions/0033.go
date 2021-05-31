package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CanCancel(num, den int) bool {
	actual := float64(num) / float64(den)
	numSplit := strings.Split(strconv.Itoa(num), "")
	denSplit := strings.Split(strconv.Itoa(den), "")
	if numSplit[0] == denSplit[1] {
		newN, _ := strconv.Atoi(numSplit[1])
		newD, _ := strconv.Atoi(denSplit[0])
		testVal := float64(newN) / float64(newD)
		return actual == testVal
	} else if numSplit[1] == denSplit[0] {
		newN, _ := strconv.Atoi(numSplit[0])
		newD, _ := strconv.Atoi(denSplit[1])
		testVal := float64(newN) / float64(newD)
		return actual == testVal
	} else {
		return false
	}
	return true
}

type Fraction struct {
	num, den     int
	numSp, denSp []string
}

func fracGenerator() <-chan Fraction {
	fracCh := make(chan Fraction)

	go func() {
		defer close(fracCh)
		for num := 11; num < 100; num++ {
			for den := num + 1; den < 100; den++ {
				if num%10 == 0 || den%10 == 0 {
					continue
				}

				fracCh <- Fraction{
					num:   num,
					den:   den,
					numSp: strings.Split(strconv.Itoa(num), ""),
					denSp: strings.Split(strconv.Itoa(den), ""),
				}
			}
		}
	}()

	return fracCh
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())

	nums, dens := 1, 1
	for x := range fracGenerator() {
		if CanCancel(x.num, x.den) {
			nums *= x.num
			dens *= x.den
		}
	}
	fmt.Println(nums)
	fmt.Println(dens)
	//fmt.Println("Sum of products:", FindSumOfProducts())
}
