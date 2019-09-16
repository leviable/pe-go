package main

import (
	"fmt"
	"time"
)

type Point struct {
	x, y int
}

func reducePaths(points []Point, gridSize int) int {
	if points[0].x == gridSize && points[0].y == gridSize {
		return len(points)
	}
	newPoints := []Point{}
	for _, point := range points {
		if point.x+1 <= gridSize {
			newPoints = append(newPoints, Point{point.x + 1, point.y})
		}
		if point.y+1 <= gridSize {
			newPoints = append(newPoints, Point{point.x, point.y + 1})
		}
	}

	return reducePaths(newPoints, gridSize)
}

func FindPaths(gridSize int) int {
	return reducePaths([]Point{Point{0, 0}}, gridSize)
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func doit(size int) {
	defer timeIt(time.Now())
	fmt.Printf("Path count for %d is %d\n", size, FindPaths(size))
}

func main() {
	for i := 16; i < 17; i++ {
		doit(i)
	}
}
