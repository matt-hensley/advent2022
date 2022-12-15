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
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

func solve(input string, contain_only bool) int {
	pairs := strings.Split(input, "\n")
	count := 0

	for _, pair := range pairs {
		assignments := strings.Split(pair, ",")
		elf1 := strings.Split(assignments[0], "-")
		elf1Start, _ := strconv.Atoi(elf1[0])
		elf1End, _ := strconv.Atoi(elf1[1])
		elf2 := strings.Split(assignments[1], "-")
		elf2Start, _ := strconv.Atoi(elf2[0])
		elf2End, _ := strconv.Atoi(elf2[1])

		if contain_only {
			if elf1Start <= elf2Start && elf2End <= elf1End {
				count += 1
			} else if elf2Start <= elf1Start && elf1End <= elf2End {
				count += 1
			}
		} else {
			if elf1Start <= elf2Start && elf2Start <= elf1End {
				count += 1
			} else if elf1Start <= elf2End && elf2End <= elf1End {
				count += 1
			} else if elf2Start <= elf1Start && elf1Start <= elf2End {
				count += 1
			} else if elf2Start <= elf1End && elf1End <= elf2End {
				count += 1
			}
		}
	}

	return count
}

func part1(input string) int {
	return solve(input, true)
}

func part2(input string) int {
	return solve(input, false)
}
