package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
)

func main() {
	file, err := os.Open("./test.txt")
	defer file.Close()

	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	fmt.Println(challenge1(fileScanner))
}

func challenge1(fileScanner *bufio.Scanner) int {
	// load a set of three lines
	// check if the line contains numbers with symbols next to them.
	// else if check if the numbers have symbols above them or diagonally (above). Use the symbol index and the number index.
	// else if same but below
	// leave the last line & load two more and repeat until no more lines to add.
	// lineIdx := 0
	numsPattern := `[0-9]+\b\W{0}`
	symbolsPattern := `[^0-9.]`
	reNums := regexp.MustCompile(numsPattern)
	reSymb := regexp.MustCompile(symbolsPattern)

	lines := []string{}
	// partNums := []int{}

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())

		if len(lines) == 3 {
			// fmt.Println("lines", lines)
			nums := [][][]int{}
			symbs := [][][]int{}
			for _, line := range lines {
				nums = append(nums, reNums.FindAllStringIndex(line, -1))
				symbs = append(symbs, reSymb.FindAllStringIndex(line, -1))
				// fmt.Println(line)
			}
			for _, num := range nums {
				// fmt.Println("num", num, i)
				if len(num) > 0 {
					for i, coordinates := range num {
						startNum := coordinates[0]
						endNum := coordinates[1]
						fmt.Println(lines[i][startNum:endNum])

					}
				}
				// for _, symb := range symbs {
				// 	startSymb := symb[0]
				// 	endSymb := symb[1]
				//
				// 	// check if the number has a symbol before or after
				// 	if endNum == startSymb || startNum == endSymb {
				// 		n, _ := strconv.Atoi(line[startNum:endNum])
				// 		partNums = append(partNums, n)
				// 	}
				// }
			}
			// fmt.Println("nums", nums)
			// fmt.Println("symbs", symbs)
			lines = slices.Delete(lines, 0, 2)
		}
		// lineIdx++
	}
	// fmt.Println(lines)
	// fmt.Println("parts", partNums)
	return 0
}
