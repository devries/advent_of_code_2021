package main

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day09.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	grid := parseGrid(lines)
	lowPoints := make([]utils.Point, 0)

outer:
	for k, v := range grid {
		for _, d := range utils.Directions {
			if va, ok := grid[k.Add(d)]; ok {
				if va <= v {
					continue outer
				}
			}
		}
		lowPoints = append(lowPoints, k)
	}

	sizes := make([]int, 0)
	for _, p := range lowPoints {
		size := findBasin(p, grid)
		sizes = append(sizes, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	res := 1
	for i := 0; i < 3; i++ {
		res *= sizes[i]
	}

	return res
}

func parseGrid(lines []string) map[utils.Point]int {
	r := make(map[utils.Point]int)

	for j, ln := range lines {
		for i, c := range ln {
			r[utils.Point{X: i, Y: j}] = int(c - '0')
		}
	}

	return r
}

func findBasin(p utils.Point, grid map[utils.Point]int) int {
	basin := make([]utils.Point, 0)
	height := grid[p]
	basin = append(basin, p)
	found := make(map[utils.Point]bool)

	for i := height + 1; i < 9; i++ {
		newBasin := make([]utils.Point, 0)
		for _, v := range basin {
			for _, d := range utils.Directions {
				pc := v.Add(d)
				if va, ok := grid[pc]; ok {
					if va == i && !found[pc] {
						newBasin = append(newBasin, pc)
						found[pc] = true
					}
				}
			}
		}
		basin = append(basin, newBasin...)
	}

	return len(basin)
}
