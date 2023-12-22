package main

import (
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

	s := string(b)
	spl := strings.Split(s, "\n\n")

	fmt.Println(firstChallenge(spl))
	fmt.Println(secondChallenge(spl))
}

func firstChallenge(spl []string) int {
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

	locations := []int{}

	for _, s := range seeds {
		seed, _ := strconv.Atoi(s)

		soil := checkVal(seedToSoil, seed)
		fert := checkVal(soilToFert, soil)
		water := checkVal(fertToWater, fert)
		light := checkVal(waterToLight, water)
		temp := checkVal(lightToTemp, light)
		hum := checkVal(tempToHum, temp)
		location := checkVal(humToLocation, hum)

		locations = append(locations, location)
	}

	return slices.Min(locations)
}

func checkVal(arr []string, s int) int {
	for _, line := range arr {
		r := strings.Fields(line)
		dest, _ := strconv.Atoi(r[0])
		source, _ := strconv.Atoi(r[1])
		rangeLen, _ := strconv.Atoi(r[2])

		if s >= source && s < source+rangeLen {
			// fmt.Println(r, source, rangeLen, s)
			// 53 => 53 - 50 => 52 + 3
			// mr := s - source
			// mapedVal := dest + mr
			mr := s - source
			res := dest + mr
			// fmt.Println(s, source, source+rangeLen, mr, res)
			return res
		}

	}
	return s
}

func secondChallenge(spl []string) int {
	_, after, _ := strings.Cut(spl[0], ":")
	seeds := strings.Fields(after)
	// seedToSoil := strings.Split(strings.Split(spl[1], ":")[1], "\n")[1:]
	// soilToFert := strings.Split(strings.Split(spl[2], ":")[1], "\n")[1:]
	// fertToWater := strings.Split(strings.Split(spl[3], ":")[1], "\n")[1:]
	// waterToLight := strings.Split(strings.Split(spl[4], ":")[1], "\n")[1:]
	// lightToTemp := strings.Split(strings.Split(spl[5], ":")[1], "\n")[1:]
	// tempToHum := strings.Split(strings.Split(spl[6], ":")[1], "\n")[1:]
	humToLocation := strings.Split(strings.Split(spl[7], ":")[1], "\n")[1:]
	humToLocation = humToLocation[:len(humToLocation)-1]

	locations := []int{}

	for i := 0; i < len(seeds); i += 2 {
		seed, _ := strconv.Atoi(seeds[i])
		r, _ := strconv.Atoi(seeds[i+1])
		fmt.Println(seed, r)

		// soil := checkVal(seedToSoil, seed)
		// fert := checkVal(soilToFert, soil)
		// water := checkVal(fertToWater, fert)
		// light := checkVal(waterToLight, water)
		// temp := checkVal(lightToTemp, light)
		// hum := checkVal(tempToHum, temp)
		// location := checkVal(humToLocation, hum)
		//
		// fmt.Println(seed, soil, fert, water, light, temp, hum, location)
		//
		// locations = append(locations, location)

		// TODO: Avoid nested loop
		// for j := 0; j < r; j++ {
		//
		// 	soil := checkVal(seedToSoil, seed+j)
		// 	fert := checkVal(soilToFert, soil)
		// 	water := checkVal(fertToWater, fert)
		// 	light := checkVal(waterToLight, water)
		// 	temp := checkVal(lightToTemp, light)
		// 	hum := checkVal(tempToHum, temp)
		// 	location := checkVal(humToLocation, hum)
		// 	fmt.Println(seed+j, soil, fert, water, light, temp, hum, location)
		//
		// 	locations = append(locations, location)
		//
		// }

	}

	return slices.Min(locations)
}
