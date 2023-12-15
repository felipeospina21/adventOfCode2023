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
	file, err := os.Open("./input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	fmt.Println(challenge1(fileScanner))
	fmt.Println(challenge2(fileScanner))
}

func challenge1(fileScanner *bufio.Scanner) int {
	// read file line by line
	// from each line extract the game id, subset of revealed cubes
	// - the game id is separated by :
	// - each subset is separated by ; and contains tuples of int string separated by , => ex: (2 red), (5 blue)
	// iterate subset and check if the quantity of cubes is higher than the goal
	// if higher break and continue with next line
	// else accumulate the game index
	// return the accum

	limit := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	games := []int{}

	for fileScanner.Scan() {
		check := []bool{}
		t := fileScanner.Text()
		spl := strings.Split(t, ":")
		gameId := strings.Split(spl[0], " ")[1]
		subset := strings.Split(spl[1], ";")

		for _, str := range subset {
			cubes := strings.Split(str, ",")

			for _, cube := range cubes {
				q, color := getCubeCountAndColor(cube)

				if q > limit[color] {
					check = append(check, false)
					break
				}
				check = append(check, true)

			}
		}

		if !slices.Contains(check, false) {
			id, _ := strconv.Atoi(gameId)
			games = append(games, id)
		}
	}

	sum := reduce(games, func(acc, current int) int {
		return acc + current
	}, 0)

	return sum
}

func challenge2(fileScanner *bufio.Scanner) int {
	accum := 0
	for fileScanner.Scan() {
		minVals := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		t := fileScanner.Text()
		spl := strings.Split(t, ":")
		subset := strings.Split(spl[1], ";")

		for _, str := range subset {
			cubes := strings.Split(str, ",")

			for _, cube := range cubes {
				q, color := getCubeCountAndColor(cube)

				if q > minVals[color] {
					minVals[color] = q
				}
			}
		}
		power := 1
		for _, val := range minVals {
			power *= val
		}

		accum += power
	}
	return accum
}

func getCubeCountAndColor(cube string) (int, string) {
	c := strings.Fields(cube)
	x, _ := strconv.Atoi(c[0])
	return x, c[1]
}

func reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}
