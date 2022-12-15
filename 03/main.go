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

func part1(input string) int {
	rucksacks := strings.Split(input, "\n")
	sum := 0

	for _, rucksack := range rucksacks {
		midpoint := len(rucksack) / 2
		comp1, comp2 := rucksack[:midpoint], rucksack[midpoint:]
		dupes := intersect(comp1, comp2)

		for _, dupe := range dupes {
			sum += item_priority(dupe)
		}
	}

	return sum
}

func part2(input string) int {
	rucksacks := strings.Split(input, "\n")
	sum := 0

	for i := 0; i < len(rucksacks); i = i + 3 {
		ruck1, ruck2, ruck3 := rucksacks[i], rucksacks[i+1], rucksacks[i+2]
		dupes := intersect(strings.Join(intersect(ruck1, ruck2), ""), ruck3)

		for _, dupe := range dupes {
			sum += item_priority(dupe)
		}
		// a := intersect(ruck1, ruck2)
		// b := intersect(ruck2, ruck3)
		// c := intersect(ruck1, ruck3)
	}

	return sum
}

func item_priority(item string) int {
	c := int(item[0])

	// ASCII A = 65, Z = 90
	if 65 <= c && c <= 90 {
		return (c - 64) + 26
	}

	// ASCII a = 97, z = 122
	if 97 <= c && c <= 122 {
		return c - 96
	}

	panic(c)
}

func intersect(left string, right string) []string {
	seen := make(map[rune]bool)
	dupes := make(map[rune]bool)

	for _, l := range left {
		seen[l] = true
	}

	for _, r := range right {
		if seen[r] {
			dupes[r] = true
		}
	}

	keys := make([]string, 0, len(dupes))

	for key := range dupes {
		keys = append(keys, string(key))
	}

	return keys
}
