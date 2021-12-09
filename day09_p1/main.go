package main

import (
	"fmt"
	"io"
	"os"

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
	total := 0

outer:
	for k, v := range grid {
		for _, d := range utils.Directions {
			if va, ok := grid[k.Add(d)]; ok {
				if va <= v {
					continue outer
				}
			}
		}
		total += v + 1
	}

	return total
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
