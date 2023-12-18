package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("error", err)
	}

	fmt.Println(firstChallenge(b))
}

func firstChallenge(b []byte) int {
	s := string(b)
	spl := strings.Split(s, "\n\n")

	_, after, _ := strings.Cut(spl[0], ":")
	seeds := strings.Fields(after)
	seedToSoil := strings.Split(strings.Split(spl[1], ":")[1], "\n")[1:]
	soilToFert := strings.Split(strings.Split(spl[2], ":")[1], "\n")[1:]
	fertToWater := strings.Split(strings.Split(spl[3], ":")[1], "\n")[1:]
	waterToLight := strings.Split(strings.Split(spl[4], ":")[1], "\n")[1:]
	lightToTemp := strings.Split(strings.Split(spl[5], ":")[1], "\n")[1:]
	tempToHum := strings.Split(strings.Split(spl[6], ":")[1], "\n")[1:]
	humToLocation := strings.Split(strings.Split(spl[7], ":")[1], "\n")[1:]
	humToLocation = humToLocation[:len(humToLocation)-1]
	//
	s2sMap := makeMap(seedToSoil)
	s2fMap := makeMap(soilToFert)
	f2wMap := makeMap(fertToWater)
	w2lMap := makeMap(waterToLight)
	l2tMap := makeMap(lightToTemp)
	t2hMap := makeMap(tempToHum)
	h2lMap := makeMap(humToLocation)

	locations := []int{}

	for _, s := range seeds {
		seed, _ := strconv.Atoi(s)
		soil := assignMapVal(s2sMap[seed], seed)
		fert := assignMapVal(s2fMap[soil], soil)
		water := assignMapVal(f2wMap[fert], fert)
		light := assignMapVal(w2lMap[water], water)
		temp := assignMapVal(l2tMap[light], light)
		hum := assignMapVal(t2hMap[temp], temp)
		location := assignMapVal(h2lMap[hum], hum)
		// fmt.Println("soil", h2lMap)
		// fmt.Sprint(hum)
		//
		locations = append(locations, location)
	}

	return slices.Min(locations)
}

func assignMapVal(dest int, source int) int {
	var x int
	if dest != 0 {
		x = dest
	} else {
		x = source
	}

	return x
}

func makeMap(arr []string) map[int]int {
	m := make(map[int]int)
	for _, line := range arr {
		r := strings.Fields(line)
		source, _ := strconv.Atoi(r[0])
		dest, _ := strconv.Atoi(r[1])
		rangeLen, _ := strconv.Atoi(r[2])

		for i := 0; i < rangeLen; i++ {
			m[dest+i] = source + i
		}
	}

	return m
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
