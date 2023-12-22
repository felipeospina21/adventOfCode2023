package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const file = "./input.txt"

func main() {
	file, err := os.Open(file)
	defer file.Close()

	if err != nil {
		log.Fatal("Error: ", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	coordinates := make(map[string][]string)
	var directions string
	currCoordinate := "AAA"

	idx := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if idx == 0 {
			directions = line
		} else if idx > 1 {

			x := strings.Fields(line)
			key := x[0]
			val := []string{x[2][1:4], x[3][:3]}

			coordinates[key] = val
		}
		idx++
	}

	fmt.Println(challenge1(directions, coordinates, currCoordinate))
	// fmt.Println(challenge2(fileScanner))
}

func challenge1(directions string, coordinates map[string][]string, currCoordinate string) int {
	var count int
	dirIdx := 0

	for {

		dir := string(directions[dirIdx])

		if dir == "L" {
			currCoordinate = coordinates[currCoordinate][0]
		}

		if dir == "R" {
			currCoordinate = coordinates[currCoordinate][1]
		}

		dirIdx++
		if dirIdx == len(directions) {
			dirIdx = 0
		}

		count++
		if currCoordinate == "ZZZ" {
			break
		}

	}
	return count
}
