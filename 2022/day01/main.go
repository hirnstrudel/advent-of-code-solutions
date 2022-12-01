package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func findBiggestRation(input string, num uint) int {
	lines := strings.Split(input, "\n")
	var biggestRations = make([]int, num)
	var currentRation int

	for i, line := range lines {
		if line == "" {
			continue
		}

		intVal, _ := strconv.Atoi(line)
		currentRation += intVal

		if i == len(lines)-1 || lines[i+1] == "" {
			for pos, v := range biggestRations {
				if currentRation > v {
					previousRation := biggestRations[pos]
					biggestRations[pos] = currentRation
					currentRation = previousRation
				}
			}
			currentRation = 0
		}

	}

	var total int
	for _, v := range biggestRations {
		total += v
	}

	return total
}

func part1(input string) int {
	return findBiggestRation(input, 1)
}

func part2(input string) int {
	return findBiggestRation(input, 3)
}

func main() {
	fmt.Println("Solution 1: ", part1(input))
	fmt.Println("Solution 2: ", part2(input))
}
