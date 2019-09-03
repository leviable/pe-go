package main

import "errors"

var GridIndexError = errors.New("Grid index out of bounds")

func SumHorizontal(x, y, sumCount int, grid *[][]int) (sum int, err error) {
	for i := 0; i < sumCount; i++ {
		if x+i >= len(*grid) || x < 0 || y < 0 {
			return 0, GridIndexError
		}
		val := (*grid)[y][x+i]
		sum += val
	}
	return sum, nil
}

func SumVertical(x, y, sumCount int, grid *[][]int) (sum int, err error) {
	for i := 0; i < sumCount; i++ {
		if y+i >= len(*grid) || y < 0 || x < 0 {
			return 0, GridIndexError
		}
		val := (*grid)[y+i][x]
		sum += val
	}
	return sum, nil
}

func SumDiagRight(x, y, sumCount int, grid *[][]int) (sum int, err error) {
	for i := 0; i < sumCount; i++ {
		if x+i >= len(*grid) || y+i >= len(*grid) || y < 0 || x < 0 {
			return 0, GridIndexError
		}
		val := (*grid)[y+i][x+i]
		sum += val
	}
	return sum, nil
}

func SumDiagLeft(x, y, sumCount int, grid *[][]int) (sum int, err error) {
	for i := 0; i < sumCount; i++ {
		if x-i < 0 || y+i >= len(*grid) || y < 0 || x < 0 {
			return 0, GridIndexError
		}
		val := (*grid)[y+i][x-i]
		sum += val
	}
	return sum, nil
}
