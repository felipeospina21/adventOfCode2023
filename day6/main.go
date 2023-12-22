package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	t := string(b)
	spl := strings.Split(t, "\n")

	times := strings.Fields(spl[0])[1:]
	distances := strings.Fields(spl[1])[1:]

	fmt.Println(challenge1(times, distances))
	fmt.Println(challenge2(times, distances))
}

func challenge1(t []string, d []string) int {
	res := 1
	for i, s := range t {
		count := 0
		time, _ := strconv.Atoi(s)
		distance, _ := strconv.Atoi(d[i])
		// holdTime * 1 = speed
		// v = d/t
		// d = (rT - hT) * (hT)
		for j := 0; j < time; j++ {
			d := (time - j) * j
			if d > distance {
				count++
			}
		}

		res *= count
	}
	return res
}

func challenge2(t []string, d []string) int {
	newT, _ := strconv.Atoi(strings.Join(t, ""))
	newD, _ := strconv.Atoi(strings.Join(d, ""))
	res := 1
	count := 0

	for j := 0; j < newT; j++ {
		d := (newT - j) * j
		if d > newD {
			count++
		}
	}

	res *= count
	return res
}
