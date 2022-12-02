package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Solution 1: ", part1(input))
	fmt.Println("Solution 2: ", part2(input))
}

func part1(input string) int {
	return findBiggestRations(input, 1)
}

func part2(input string) int {
	return findBiggestRations(input, 3)
}

func findBiggestRations(input string, num uint) int {
	rations := splitIntoRations(input)
	biggestRations := make([]int, num)
	var currentRation int

	for _, ration := range rations {
		currentRation = addSnacksInRation(ration)
		biggestRations = insertRationIntoBiggestRations(currentRation, biggestRations)
	}

	return addSnacksInRation(biggestRations)
}

func splitIntoRations(input string) [][]int {
	var rations [][]int
	lines := strings.Split(input, "\n")
	var currentRation []int

	for i, line := range lines {
		if line == "" {
			rations = append(rations, currentRation)
			currentRation = []int{}
			continue
		}

		intValue, _ := strconv.Atoi(line)
		currentRation = append(currentRation, intValue)

		if i == len(lines)-1 {
			rations = append(rations, currentRation)
		}
	}

	return rations
}

func addSnacksInRation(ration []int) int {
	var total int
	for _, snack := range ration {
		total += snack
	}

	return total
}

func insertRationIntoBiggestRations(ration int, biggestRations []int) []int {
	for i, v := range biggestRations {
		if ration > v {
			prev := biggestRations[i]
			biggestRations[i] = ration
			ration = prev
		}
	}

	return biggestRations
}
