package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1: ", part1(input))
  fmt.Println("Part 2: ", part2(input))
}

func part1(input string) int {
  return solve(input, 4)
}

func part2(input string) int {
  return solve(input, 14)
}

func solve(input string, window int) int {
	for i := 0; i < len(input); i++ {
		seen := map[byte]bool{}

    for j := 0; j < window; j++ {
			seen[input[i + j]] = true
		}

    if len(seen) == window {
			return i + window
		}
	}

  panic("unique pattern not found")
}
