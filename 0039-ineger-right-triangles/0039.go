package main

import (
	"context"
	"fmt"
	"time"
)

// If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.
// {20,48,52}, {24,45,51}, {30,40,50}
// For which value of p ≤ 1000, is the number of solutions maximised?

type Triangle struct {
	a, b, c int
}

func solutionGenerator(ctx context.Context, perim int) <-chan Triangle {
	solCh := make(chan Triangle, 0)

	go func() {
		defer close(solCh)
		for c := perim - 2; c >= 2; c -= 1 {
			for b := c - 1; b > 0; b -= 1 {
				for a := b; a > 0; a -= 1 {
					if a+b+c != perim {
						continue
					}
					select {
					case <-ctx.Done():
					case solCh <- Triangle{a, b, c}:
					}
				}
			}
		}
		return
	}()

	return solCh
}

func solutionGenerator2(ctx context.Context, perim int) <-chan Triangle {
	solCh := make(chan Triangle, 0)

	go func() {
		defer close(solCh)
		for c := 2; c < perim-2; c++ {
			for b := 1; b < c; b++ {
				for a := 1; a <= perim-c-b && a <= b; a++ {
					t := Triangle{a, b, c}
					if !IsRightTriangle(t) {
						continue
					}
					select {
					case <-ctx.Done():
					case solCh <- t:
					}
				}
			}
		}
		return
	}()

	return solCh
}

func GetPerimeterSolutions(perim int) []Triangle {
	ctx, _ := context.WithCancel(context.Background())
	solutions := make([]Triangle, 0)

	for solution := range solutionGenerator(ctx, perim) {
		if IsRightTriangle(solution) {
			solutions = append(solutions, solution)
		}
	}

	return solutions
}

func IsRightTriangle(tri Triangle) bool {
	return tri.c*tri.c == (tri.a*tri.a)+(tri.b*tri.b)
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())

	ctx := context.Background()
	solutions := make(map[int][]Triangle)

	for t := range solutionGenerator2(ctx, 1000) {
		p := t.a + t.b + t.c
		solutions[p] = append(solutions[p], t)
	}

	max := 0
	for k, v := range solutions {
		if len(v) > max {
			max = len(v)
			fmt.Println("New max: ", k)
		}
	}
	//fmt.Println("Sum of first 11 truncatable primes: ", foo(11))
}
