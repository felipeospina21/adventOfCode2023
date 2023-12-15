package main

import (
	"advent2023/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal("Error reading file, ", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	// fmt.Println(firstChallenge(fileScanner))
	fmt.Println(secondChallenge(fileScanner))
}

func firstChallenge(fileScanner *bufio.Scanner) int {
	// read each line
	// create two slices (separated by | in the line)
	// one with wining nums
	// other with nums I have
	// iterate nums I have slice and check if num is in wining slice
	// if is in wining silce && points <= 1
	// points ++
	// if is in wining slice && poinst > 1
	// points *= points

	totalPoints := []int{}

	for fileScanner.Scan() {
		points := 0
		line := fileScanner.Text()

		nums := strings.Split(line, ":")[1]
		splNums := strings.Split(nums, "|")

		winningNums := strings.Split(strings.TrimSpace(splNums[0]), " ")
		haveNums := strings.Split(strings.TrimSpace(splNums[1]), " ")

		for _, num := range haveNums {
			isWinningNum := slices.Contains(winningNums, num)

			if isWinningNum && points <= 1 && num != "" {
				points++
			} else if isWinningNum && points > 1 && num != "" {
				points *= 2
			}
		}
		totalPoints = append(totalPoints, points)

	}
	return utils.Reduce(totalPoints, func(acc, current int) int {
		return acc + current
	}, 0)
}

func secondChallenge(fileScanner *bufio.Scanner) int {
	// create map[int]int to hold for each card num the instances
	// create var card := 1
	// read each line
	// create points var
	// create two slices (separated by | in the line)
	// one with wining nums
	// other with nums I have
	// iterate nums I have slice and check if num is in wining slice
	// points ++
	// iterate for range points to add instances to the map
	// map[card+1] ++
	// if map[card] > 0
	// iterate for range map[card]
	// map[card+1] ++
	// card ++

	cardInstances := make(map[int]int)
	card := 1

	for fileScanner.Scan() {
		points := 0
		line := fileScanner.Text()

		nums := strings.Split(line, ":")[1]
		splNums := strings.Split(nums, "|")

		winningNums := strings.Split(strings.TrimSpace(splNums[0]), " ")
		haveNums := strings.Split(strings.TrimSpace(splNums[1]), " ")

		for _, num := range haveNums {
			isWinningNum := slices.Contains(winningNums, num)

			if isWinningNum && num != "" {
				points++
			}
		}

		for i := 0; i < points; i++ {
			cardInstances[card+1+i]++
			if cardInstances[card] > 0 {
				cardInstances[card+1+i] += cardInstances[card]
			}
		}

		cardInstances[card]++

		card++
	}
	result := 0
	for _, instances := range cardInstances {
		result += instances
	}

	return result
}

// card 1 => 4 => 2,3,4,5 {2:1, 3:1, 4:1, 5:1}

// card 2 => 2 => 3,4     {2:1, 3:2, 4:2, 5:1}
// COPY 2 => 2 => 3,4     {2:1, 3:3, 4:3, 5:1}

// card 3 => 2 => 4,5     {2:1, 3:3, 4:4, 5:2}
// card 3 => 2 => 4,5     {2:1, 3:3, 4:5, 5:3}
// card 3 => 2 => 4,5     {2:1, 3:3, 4:6, 5:4}
// card 3 => 2 => 4,5     {2:1, 3:3, 4:7, 5:5}

// card 4 => 1 => 4,5     {2:1, 3:3, 4:7, 5:6}
// card 4 => 1 => 4,5     {2:1, 3:3, 4:7, 5:7}
// card 4 => 1 => 4,5     {2:1, 3:3, 4:7, 5:8}
// card 4 => 1 => 4,5     {2:1, 3:3, 4:7, 5:9}
// card 4 => 1 => 4,5     {2:1, 3:3, 4:7, 5:10}
// card 4 => 1 => 4,5     {2:1, 3:3, 4:7, 5:11}
// card 4 => 1 => 4,5     {2:1, 3:3, 4:7, 5:12}
// card 4 => 1 => 4,5     {2:1, 3:3, 4:7, 5:13}
