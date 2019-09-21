package main

import (
	"fmt"
	"time"
)

var Triangle = [][]int{
	{75},
	{95, 64},
	{17, 47, 82},
	{18, 35, 87, 10},
	{20, 4, 82, 47, 65},
	{19, 1, 23, 75, 3, 34},
	{88, 2, 77, 73, 7, 63, 67},
	{99, 65, 4, 28, 6, 16, 70, 92},
	{41, 41, 26, 56, 83, 40, 80, 70, 33},
	{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
	{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
	{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
	{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
	{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
	{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
}

func ReducePaths(triangle [][]int, x, y int) (sum int) {
	curVal := triangle[y][x]
	if y == 0 {
		return curVal
	} else if x == 0 {
		return curVal + ReducePaths(triangle, x, y-1)
	} else if x == len(triangle[y])-1 {
		return curVal + ReducePaths(triangle, x-1, y-1)
	} else {
		left := ReducePaths(triangle, x-1, y-1)
		right := ReducePaths(triangle, x, y-1)
		if left >= right {
			return curVal + left
		} else {
			return curVal + right
		}
	}
}

func FindGreatestPath(triangle [][]int) (pathSum int) {
	y := len(triangle) - 1

	for i := 0; i < len(triangle[y]); i++ {
		pathVal := ReducePaths(triangle, i, y)
		if pathVal > pathSum {
			pathSum = pathVal
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
	fmt.Println("Max path val: ", FindGreatestPath(Triangle))
}
