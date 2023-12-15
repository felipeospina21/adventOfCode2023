package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(reduceValues())
}

func reduceValues() int {
	// readFile, err := os.Open("./sample.txt")
	readFile, err := os.Open("./input.txt")
	defer readFile.Close()

	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return secondPuzzleCalc(fileScanner)
	// return firstPuzzleCalc(fileScanner)
}

func firstPuzzleCalc(fileScanner *bufio.Scanner) int {
	accum := 0
	for fileScanner.Scan() {
		t := fileScanner.Text()
		re := regexp.MustCompile(`[1-9]`)
		m := re.FindAllString(t, -1)

		a, _ := strconv.Atoi(m[0])
		b, _ := strconv.Atoi(m[len(m)-1])
		accum += a*10 + b

	}

	return accum
}

func secondPuzzleCalc(fileScanner *bufio.Scanner) int {
	accum := 0
	adapter := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for fileScanner.Scan() {
		var first, last int
		t := fileScanner.Text()
		evalString := ""
		patternNums := `([1-9])`
		patternStrings := `(one|two|three|four|five|six|seven|eight|nine)`

		re := regexp.MustCompile(patternNums + "|" + patternStrings)

		for _, r := range t {
			evalString += string(r)
			m := re.FindString(evalString)
			if m != "" {
				if adapter[m] != 0 {
					first = adapter[m]
				} else {
					x, _ := strconv.Atoi(m)
					first = x
				}
				break
			}
		}
		evalString = ""

		for i := len(t) - 1; i >= 0; i-- {
			evalString = string(t[i]) + evalString
			m := re.FindString(evalString)
			if m != "" {
				if adapter[m] != 0 {
					last = adapter[m]
				} else {
					x, _ := strconv.Atoi(m)
					last = x
				}
				break
			}

		}
		accum += first*10 + last
	}

	return accum
}

func reverseString(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
