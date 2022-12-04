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

  for _, rucksack := range rucksacks {
    midpoint := len(rucksack) / 2
    comp1, comp2 = rucksack[:midpoint], rucksack[midpoint:]
    
    // items := strings.Split(rucksack, "")
    
  }
}

func part2(input string) int {
}

func item_priority(item string) int {
  c := int(item[0])

  // ASCII A = 65, Z = 90
  if 65 <= c && c <= 90 {
    return c - 64
  }
}
