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
	fmt.Println("Part 2: \n", part2(input))
}

func part1(input string) int {
	register := generate_register(input)
	sum := 0

	for cycle := 0; cycle < len(register); cycle++ {
		if cycle%40 == 20 {
			sum += cycle * register[cycle-1]
		}
	}

	return sum
}

func part2(input string) string {
	register := generate_register(input)
	str := ""

	for cycle, X := range register {
		col := cycle % 40

		if col == X-1 || col == X || col == X+1 {
			str += fmt.Sprint("#")
		} else {
			str += fmt.Sprint(".")
		}

		if col == 39 {
			str += fmt.Sprint("\n")
		}
	}

	return str
}

func generate_register(input string) []int {
	lines := strings.Split(input, "\n")
	register := []int{1}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		last := register[len(register)-1]

		switch parts[0] {
		case "noop":
			register = append(register, last)
		case "addx":
			val, _ := strconv.Atoi(parts[1])
			register = append(register, last)
			register = append(register, last+val)
		}
	}

	return register
}
