package main

import (
	_ "embed"
	"fmt"
  "math"
	"strings"
)

//go:embed input.txt
var input string

type Point struct{ x, y int }

func (p *Point) equals(comp Point) bool {
	return p.x == comp.x && p.y == comp.y
}

func main() {
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

func part1(input string) int {
  grid, start, end := build_grid(input)
  return search(grid, start, end)
}

func part2(input string) int {
	return -1
}

func search(grid [][]int, start Point, end Point) int {
  shortest := math.MaxInt
  queue := []Point{start}

  for len(queue) > 0 {
    current := queue[0]
    queue = queue[1:]
    
    if current.equals(end) {
      
    }

    for _, neighbor := range neighbors(grid, current) {
      
    }
    
  }

  return shortest
}

func build_grid(input string) (grid [][]int, start Point, end Point) {
	grid = [][]int{}

	for y, line := range strings.Split(input, "\n") {
		row := []int{}

		for x, c := range line {
			if c == 'S' {
				start = Point{x, y}
				c = 'a'
			} else if c == 'E' {
				end = Point{x, y}
				c = 'z'
			}

			// ASCII a = 97, z = 122
			row = append(row, int(c)-96)
		}

		grid = append(grid, row)
	}

	return grid, start, end
}

func neighbors(grid [][]int, p Point) []Point {
  neighbors := []Point{}

  for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		neighbor := Point{p.x + dir.x, p.y + dir.y}
    
		if neighbor.x >= 0 && neighbor.x < len(grid) && neighbor.y >= 0 && neighbor.y < len(grid[0]) {
			if abs(int(grid[neighbor.x][neighbor.y])-int(grid[p.x][p.y])) == 1 {
				neighbors = append(neighbors, neighbor)
			}
		}
	}

  return neighbors
}

func main2() {
	grid := [][]byte{
		{'S', 'a', 'b', 'q', 'p', 'o', 'n', 'm'},
		{'a', 'b', 'c', 'r', 'y', 'x', 'x', 'l'},
		{'a', 'c', 'c', 's', 'z', 'E', 'x', 'k'},
		{'a', 'c', 'c', 't', 'u', 'v', 'w', 'j'},
		{'a', 'b', 'd', 'e', 'f', 'g', 'h', 'i'},
	}
	start := Point{0, 0}
	end := Point{5, 7}
	shortestPath := findShortestPath(grid, start, end)
	fmt.Println("Shortest path:", shortestPath)
}

func findShortestPath(grid [][]byte, start, end Point) int {
	queue := []Point{start}
	visited := map[Point]bool{start: true}
	distances := map[Point]int{start: 0}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return distances[current]
		}

		neighbors := getNeighbors(grid, current)

    for _, neighbor := range neighbors {
			if _, ok := visited[neighbor]; !ok {
				queue = append(queue, neighbor)
				visited[neighbor] = true
				distances[neighbor] = distances[current] + 1
			}
		}
	}
  
	return math.MaxInt32
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
