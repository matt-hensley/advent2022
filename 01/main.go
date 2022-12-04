package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	var max, _ = solve(input)
	return max
}

func part2(input string) int {
	var _, last3sum = solve(input)
	return last3sum
}

func solve(input string) (int, int) {
	var parts []string = strings.Split(input, "\n\n")
	var sums []int

	for _, part := range parts {
		calories := strings.Split(part, "\n")
		sum := 0

		for _, calorie := range calories {
			val, err := strconv.Atoi(calorie)

			if err != nil {
				panic("bad number")
			}

			sum += val
		}

		sums = append(sums, sum)
	}

	sort.Ints(sums)
	max := sums[len(sums)-1]
	last3 := sums[len(sums)-3:]

	last3sum := 0
	for _, x := range last3 {
		last3sum += x
	}

	return max, last3sum
}
