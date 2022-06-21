package main

import (
	"fmt"
	"time"
)

func PE0000() {

	return
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	defer timeIt(time.Now())
	fmt.Println("0000: ", PE0000())
}
