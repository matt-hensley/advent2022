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

type File struct {
	name string
	size int
}

type Directory struct {
	name     string
	parent   *Directory
	children map[string]*Directory
	files    map[string]*File
}

func (dir *Directory) size() int {
	size := 0

	for _, child := range dir.children {
		size += child.size()
	}

	for _, file := range dir.files {
		size += file.size
	}

	return size
}

func newDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		name,
		parent,
		make(map[string]*Directory),
		make(map[string]*File),
	}
}

func newFile(name string, size int) *File {
	return &File{name, size}
}

func main() {
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

func part1(input string) int {
  dirs := build_tree(input)
  sum := 0

	for _, dir := range dirs {
		size := dir.size()

		if size <= 100000 {
			sum += size
		}
	}

	return sum
}


func part2(input string) int {
  max := 70000000
  target := 30000000
  dirs := build_tree(input)
  current := max - dirs[0].size()

  sort.SliceStable(dirs, func (l int, r int) bool {
    return dirs[l].size() < dirs[r].size()
  })

  for _, dir := range dirs {
    size := dir.size()

    if current + size >= target {
      return size
    }
  }
  
	return -1
}


func build_tree(input string) []*Directory {
	var root *Directory
	var current *Directory
	dirs := make([]*Directory, 0)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, " ")

		switch parts[0] {
		case "$":
			// command
			switch parts[1] {
			case "cd":
				// traverse
				name := parts[2]

				if name == ".." {
					// should check for .. sanity
					current = current.parent
					continue
				}

				if root == nil {
          // only create root dir
          // other dirs are seen by ls command first
					dir := newDirectory(name, current)
					dirs = append(dirs, dir)
					root = dir
					current = dir
          continue
				}

        current = current.children[name]
			case "ls":
				// list
			default:
				panic("no command default")
			}
		case "dir":
			name := parts[1]

			if _, exists := current.children[name]; exists {
				continue
			}

      dir := newDirectory(name, current)
      dirs = append(dirs, dir)
			current.children[name] = dir
		default:
			// file?
			size, err := strconv.Atoi(parts[0])

			if err != nil {
				panic(err)
			}

			name := parts[1]
			current.files[name] = newFile(name, size)
		}
	}

	return dirs
}
