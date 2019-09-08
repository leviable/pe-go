package main

import (
	"fmt"
	"time"
)

type LongestSequence struct {
	Value  int
	Length int
}

func NextCollatzNum(num int) (next int) {
	if num%2 == 0 {
		return num / 2
	} else {
		return 3*num + 1
	}
}

func GetCollatzSeq(num int) (seq []int) {
	if num == 1 {
		return []int{}
	}
	next := NextCollatzNum(num)
	return append(GetCollatzSeq(next), next)
}

func GetLongestCollatzSeq(num int) (longest LongestSequence) {
	for i := 1; i < num+1; i++ {
		seq := GetCollatzSeq(i)
		newSeq := LongestSequence{i, len(seq) + 1}
		if newSeq.Length > longest.Length {
			longest = newSeq
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
	fmt.Println("Longest sequence is ", GetLongestCollatzSeq(1000000))
}
