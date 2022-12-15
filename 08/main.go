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

func part1(input string) int {
	grid := build_grid(input)
	height := len(grid)
	width := len(grid[0])
	visible := 0

	for y, row := range grid {
		if y == 0 || y == height-1 {
			// first and last column
			continue
		}

		for x := range row {
			if x == 0 || x == width-1 {
				// first and last row
				continue
			}

			left := max(grid[y][:x])
			right := max(grid[y][x+1:])
			current := grid[y][x]

			if left < current || right < current {
				visible += 1
				continue
			}

			col := column(grid, x)
			up := max(col[:y])
			down := max(col[y+1:])

			if up < current || down < current {
				visible += 1
				continue
			}
		}
	}

	visible += height * 2
	visible += width * 2
	visible -= 4 // 4 corners

	return visible
}

func part2(input string) int {
	grid := build_grid(input)
	max := 0

	for y, row := range grid {
		for x := range row {
			left := grid[y][:x]
			right := grid[y][x+1:]
			current := grid[y][x]

			scores := make([]int, 0)
			scores = append(scores, scenic_score(current, left, true))
			scores = append(scores, scenic_score(current, right, false))

			col := column(grid, x)
			up := col[:y]
			down := col[y+1:]
			scores = append(scores, scenic_score(current, up, true))
			scores = append(scores, scenic_score(current, down, false))
			score := mult(scores)

			if score > max {
				max = score
			}
		}
	}

	//fmt.Println(scores)

	return max
}

func mult(values []int) int {
	acc := 1

	for _, value := range values {
		acc *= value
	}

	return acc
}

func max(values []int) int {
	max := values[0]

	for _, value := range values {
		if max < value {
			max = value
		}
	}

	return max
}

func column(grid [][]int, x int) []int {
	col := make([]int, 0)

	for y := range grid {
		col = append(col, grid[y][x])
	}

	return col
}

func scenic_score(target int, trees []int, reverse bool) int {
	count := 0
	if reverse == false {
		for i := 0; i < len(trees); i++ {
			count++
			if trees[i] >= target {
				break
			}
		}
	} else {
		for i := len(trees) - 1; 0 <= i; i-- {
			count++
			if trees[i] >= target {
				break
			}
		}
	}

	return count
}

func build_grid(input string) [][]int {
	grid := [][]int{}

	for _, line := range strings.Split(input, "\n") {
		row := []int{}

		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			row = append(row, n)
		}

		grid = append(grid, row)
	}

	return grid
}
