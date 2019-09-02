package main

func IsTriplet(a, b, c int) bool {
	if a > b || b > c {
		return false
	}
	return a*a+b*b == c*c
}
