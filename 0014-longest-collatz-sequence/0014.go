package main

import (
	"fmt"
	"time"
)

var seqMap = map[int][]int{
	1: []int{},
}

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

func GetCollatzSeq(num int) []int {
	next := NextCollatzNum(num)
	if _, ok := seqMap[next]; !ok {
		seq := GetCollatzSeq(next)
		seqMap[next] = seq
	}
	seq := seqMap[next]
	return append(seq, next)
}

func GetLongestCollatzSeq(num int) (longest LongestSequence) {
	for i := num; i > 0; i-- {
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
	//fmt.Println("Longest sequence is ", GetCollatzSeq(13))
}
