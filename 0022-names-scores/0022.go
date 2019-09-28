package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
	"unicode"
)

func LoadNames() ([]string, error) {
	file, err := os.Open("p022_names.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	f := func(c rune) bool {
		return unicode.IsPunct(c)
	}
	return strings.FieldsFunc(lines[0], f), nil
}

func NameValue(name string) (sum int) {
	for _, c := range name {
		sum += (int(c) - 64)
	}
	return
}

func GetNameScore(name string, position int) (score int) {
	return position * NameValue(name)
}

func GetNameScoreTotal(names []string) (score int) {
	sort.Strings(names)
	for i, name := range names {
		score += GetNameScore(name, i+1)
	}
	return
}

func timeIt(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Completed in %s\n", elapsed)
}

func main() {
	names, err := LoadNames()

	if err != nil {
		panic(err)
	}

	defer timeIt(time.Now())
	fmt.Println("Total name score is", GetNameScoreTotal(names))
}
