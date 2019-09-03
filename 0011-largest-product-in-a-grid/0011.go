package main

func SumHorizontal(x, y, sumCount int, grid *[][]int) (sum int) {
	for i := 0; i < sumCount; i++ {
		val := (*grid)[y][x+i]
		sum += val
	}
	return sum
}

func SumVertical(x, y, sumCount int, grid *[][]int) (sum int) {
	for i := 0; i < sumCount; i++ {
		val := (*grid)[y+i][x]
		sum += val
	}
	return sum
}

func SumDiagRight(x, y, sumCount int, grid *[][]int) (sum int) {
	for i := 0; i < sumCount; i++ {
		val := (*grid)[y+i][x+i]
		sum += val
	}
	return sum
}

func SumDiagLeft(x, y, sumCount int, grid *[][]int) (sum int) {
	for i := 0; i < sumCount; i++ {
		val := (*grid)[y+i][x-i]
		sum += val
	}
	return sum
}
