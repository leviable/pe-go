package main

import "fmt"

func SumHorizontal(x, y, sumCount int, grid *[][]int) (sum int) {
	for i := 0; i < sumCount; i++ {
		val := (*grid)[x][y+i]
		fmt.Printf("Val is %d\n", val)
		sum += val
	}
	return sum
}
