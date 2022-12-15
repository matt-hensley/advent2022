package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

// rock, paper, scissors
const (
	Rock     int = 1
	Paper        = 2
	Scissors     = 3
)
const (
	Lose int = 0
	Draw     = 3
	Win      = 6
)

var opp_key = map[string]int{"A": Rock, "B": Paper, "C": Scissors}
var mine_key = map[string]int{"X": Rock, "Y": Paper, "Z": Scissors}
var outcomes = map[string]int{"X": Lose, "Y": Draw, "Z": Win}

func part1(input string) int {
	score := 0
	plays := strings.Split(input, "\n")

	for _, play := range plays {
		parts := strings.Split(play, " ")
		opp_encoded, mine := parts[0], parts[1]
		shape := mine_key[mine]
		round := round_score(opp_key[opp_encoded], shape)
		score += shape + round
	}

	return score
}

func part2(input string) int {
	score := 0
	plays := strings.Split(input, "\n")

	for _, play := range plays {
		parts := strings.Split(play, " ")
		opp_encoded, outcome := parts[0], parts[1]
		round := outcomes[outcome]
		shape := pick_shape(opp_key[opp_encoded], round)
		score += shape + round
	}

	return score
}

func round_score(opp int, mine int) int {
	if opp == Rock && mine == Scissors {
		return Lose
	}
	if opp == Paper && mine == Rock {
		return Lose
	}
	if opp == Scissors && mine == Paper {
		return Lose
	}
	if mine == Rock && opp == Scissors {
		return Win
	}
	if mine == Paper && opp == Rock {
		return Win
	}
	if mine == Scissors && opp == Paper {
		return Win
	}

	// default to tie
	return Draw
}

func pick_shape(opp int, outcome int) int {
	if opp == Rock {
		if outcome == Lose {
			return Scissors
		}
		if outcome == Win {
			return Paper
		}
	}

	if opp == Paper {
		if outcome == Lose {
			return Rock
		}
		if outcome == Win {
			return Scissors
		}
	}

	if opp == Scissors {
		if outcome == Lose {
			return Paper
		}
		if outcome == Win {
			return Rock
		}
	}

	// return same shape for Draw
	return opp
}
