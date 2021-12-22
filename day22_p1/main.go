package main

import (
	"fmt"
	"io"
	"os"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day22.txt")
	utils.Check(err, "error opening input file")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	grid := parseInput(lines)

	total := 0
	for _, v := range grid {
		if v {
			total++
		}
	}

	return total
}

type Point struct {
	X int
	Y int
	Z int
}

func parseInput(lines []string) map[Point]bool {
	grid := make(map[Point]bool)

	for _, line := range lines {
		var low Point
		var high Point
		var command string
		_, err := fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &command, &low.X, &high.X, &low.Y, &high.Y, &low.Z, &high.Z)
		utils.Check(err, "error parsing line")

		if low.X >= -50 && high.X <= 50 && low.Y >= -50 && high.Y <= 50 && low.Z >= -50 && high.Z <= 50 {
			for k := low.Z; k <= high.Z; k++ {
				for j := low.Y; j <= high.Y; j++ {
					for i := low.X; i <= high.X; i++ {
						if command == "on" {
							grid[Point{i, j, k}] = true
						} else {
							grid[Point{i, j, k}] = false
						}
					}
				}
			}
		}
	}

	return grid
}
