package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func GenPermutations(orig []int) (perms []string) {
	for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
		perm := getPerm(orig, p)
		// from https://stackoverflow.com/questions/37532255
		permStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(perm)), ""), "[]")
		perms = append(perms, permStr)
	}

	return
}

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	orig := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// nextPerm, getPerm, and GenPerm for loop from https://stackoverflow.com/questions/30226438
	perms := GenPermutations(orig)
	sort.Strings(perms)
	fmt.Println("Target permutation is", perms[999999])
}
