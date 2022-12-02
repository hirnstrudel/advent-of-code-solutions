package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

type Symbol uint
type Strategy uint
type Outcome uint

const (
	Loose Outcome = iota
	Draw
	Win
)

const (
	SYMBOL Strategy = iota
	OUTCOME
)

const (
	Rock Symbol = iota
	Paper
	Scissors
)

var symbolMap = map[string]Symbol{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var outcomeMap = map[string]Outcome{
	"X": Loose,
	"Y": Draw,
	"Z": Win,
}

var symbolValue = map[Symbol]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

func main() {
	fmt.Println("Solution 1: ", part1(input))
	fmt.Println("Solution 2: ", part2(input))
}

func part1(input string) int {
	return calculateScorePart(input, SYMBOL)
}

func part2(input string) int {
	return calculateScorePart(input, OUTCOME)
}

func calculateScorePart(input string, strat Strategy) int {
	var score int

	if strat == SYMBOL {
		for i := 0; i < len(input); i += 4 {
			score += scoreBySymbol(symbolMap[string(input[i])], symbolMap[string(input[i+2])])
		}
	} else {
		for i := 0; i < len(input); i += 4 {
			score += scoreByOutcome(symbolMap[string(input[i])], outcomeMap[string(input[i+2])])
		}
	}
	return score
}

func scoreBySymbol(oponentSymbol Symbol, mySymbol Symbol) int {
	value := symbolValue[mySymbol]
	if oponentSymbol == mySymbol {
		return 3 + value
	} else if (oponentSymbol+1)%3 == mySymbol {
		return 6 + value
	} else {
		return value
	}
}

func scoreByOutcome(oponentSymbol Symbol, outcome Outcome) int {
	if outcome == Loose {
		return symbolValue[(oponentSymbol+2)%3]
	} else if outcome == Draw {
		return symbolValue[oponentSymbol] + 3
	} else {
		return symbolValue[(oponentSymbol+1)%3] + 6
	}
}
