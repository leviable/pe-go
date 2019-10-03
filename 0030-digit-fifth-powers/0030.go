package main

import (
	"fmt"
	"math"
	"time"
)

func GetPowerSum(num, pow int) (sum int) {
	for i := 1; i < num; i *= 10 {
		dig := (num / i) % 10
		sum += int(math.Pow(float64(dig), float64(pow)))
	}
	return
}

func SumAllForPower(pow int) (sum int) {
	//for i := int(math.Pow10(pow - 1)); i < int(math.Pow10(pow)); i++ {
	for i := 10; i < 531441; i++ {
		if i == GetPowerSum(i, pow) {
			sum += i
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
	fmt.Println("Fifth power sum is", SumAllForPower(5))
}
