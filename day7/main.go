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

const file = "./test2.txt"

func main() {
	file, err := os.Open(file)
	defer file.Close()

	if err != nil {
		log.Fatal("Error: ", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	fmt.Println(challenge1(fileScanner))
	// fmt.Println(challenge2(fileScanner))
}

type HandType = int

const (
	highCard HandType = iota
	onePair
	twoPair
	threeOfKind
	fullHouse
	fourOfKind
	fiveOfKind
)

var cardStrength = map[string]int{"A": 12, "K": 11, "Q": 10, "J": 9, "T": 8, "9": 7, "8": 6, "7": 5, "6": 4, "5": 3, "4": 2, "3": 1, "2": 0}

//[{69K4Q 232 0} {66JQ9 398 1} {6T9AT 311 1} {53TT3 43 2} {5TTAA 647 2} {44JTT 779 2} {88223 818 2} {T4T66 496 2} {6AAA3 383 3} {J6266 762 3}]

type handT struct {
	count int
	value string
	index int
}

type orderedHand struct {
	hand     string
	bid      int
	handType HandType
}

func challenge1(fs *bufio.Scanner) int {
	orderedHands := []orderedHand{}
	lineIdx := 0
	for fs.Scan() {
		line := fs.Text()
		spl := strings.Fields(line)
		hand := spl[0]
		bid, _ := strconv.Atoi(spl[1])

		handMap := make(map[string]handT)
		count := 1

		for i, r := range hand {
			label := string(r)
			if handMap[label].count > 0 {
				handMap[label] = handT{count: handMap[label].count + 1, index: i, value: label}
				if count < handMap[label].count {
					count++
				}
			} else {
				handMap[label] = handT{count: 1, index: i, value: label}
			}
		}
		newHandType := getHandType(handMap, count)
		newHand := orderedHand{hand: hand, bid: bid, handType: newHandType}

		fmt.Println("hand", hand, orderedHands)
		if len(orderedHands) > 0 {
			handToCheck := orderedHands[lineIdx]
			// fmt.Println(hand, newHandType, handToCheck.handType, orderedHands)

			if newHandType > handToCheck.handType {
				orderedHands = slices.Insert(orderedHands, lineIdx+1, newHand)
			} else if newHandType < handToCheck.handType {
				orderedHandsLastIdx := len(orderedHands) - 1

				for orderedHandsLastIdx > 0 {
					currArrType := orderedHands[orderedHandsLastIdx].handType
					prevArrType := orderedHands[orderedHandsLastIdx-1].handType

					if newHandType > currArrType {
						orderedHands = slices.Insert(orderedHands, orderedHandsLastIdx+1, newHand)
						break
					} else if newHandType < currArrType && newHandType > prevArrType {
						orderedHands = slices.Insert(orderedHands, orderedHandsLastIdx, newHand)
						break
					} else if newHandType < currArrType && newHandType == prevArrType {
						for i := orderedHandsLastIdx - 1; i >= 0; i-- {
							if i-1 > 0 {
								fmt.Println("orde", i, orderedHands[i-1], newHandType)
								if newHandType > orderedHands[i-1].handType {
									idx := checkValue(hand, i-1, orderedHands)
									if idx == 1 {
										orderedHands = slices.Insert(orderedHands, orderedHandsLastIdx, newHand)
									} else {
										orderedHands = slices.Insert(orderedHands, orderedHandsLastIdx-1, newHand)
									}
									break
								} else if newHandType < orderedHands[i-1].handType {
									idx := checkValue(hand, i, orderedHands)
									if idx == 1 {
										orderedHands = slices.Insert(orderedHands, orderedHandsLastIdx, newHand)
									} else {
										orderedHands = slices.Insert(orderedHands, orderedHandsLastIdx-1, newHand)
									}
									break
								}
							} else {
								orderedHands = slices.Insert(orderedHands, i, newHand)
							}
						}
						break

					} else {
						orderedHandsLastIdx--
					}
				}
				if orderedHandsLastIdx == 0 {
					orderedHands = slices.Insert(orderedHands, 0, newHand)
				}
			} else {
				idx := checkValue(hand, lineIdx, orderedHands)
				orderedHands = slices.Insert(orderedHands, lineIdx+idx, newHand)
			}
			lineIdx++
		} else {
			orderedHands = append(orderedHands, newHand)
		}
	}

	var res int
	for i, oh := range orderedHands {
		res += oh.bid * (i + 1)
	}

	fmt.Println("res", orderedHands)
	return res
}

func checkValue(hand string, index int, orderedHands []orderedHand) int {
	for j, r := range hand {
		currCard := string(r)
		sliceCard := string(orderedHands[index].hand[j])
		// fmt.Println(orderedHands[index])
		fmt.Println("from check", currCard, cardStrength[currCard], sliceCard, cardStrength[sliceCard], orderedHands[index])

		if cardStrength[currCard] > cardStrength[sliceCard] {
			return 1
		} else if cardStrength[currCard] < cardStrength[sliceCard] {
			return 0
		}
	}

	return 0
}

func getHandType(handMap map[string]handT, count int) int {
	switch len(handMap) {
	case 5:
		return highCard

	case 4:
		return onePair

	case 3:
		if count == 3 {
			return threeOfKind
		}
		return twoPair

	case 2:
		return fullHouse

	case 1:
		return fourOfKind

	default:
		return fiveOfKind

	}
}

func challenge2(fs *bufio.Scanner) int {
	return 0
}
