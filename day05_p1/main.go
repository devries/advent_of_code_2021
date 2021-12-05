package main

import (
	"fmt"
	"io"
	"os"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day05.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	grid := make(map[utils.Point]int)

	for _, ln := range lines {
		p1, p2 := parseLine(ln)

		// find difference and then direction between points
		d := p2.Add(p1.Scale(-1))
		if d.X != 0 && d.Y != 0 {
			// Not horizontal or vertical
			continue
		}

		ds := direction(d)

		// Add the vents
		// This is a do ... while loop, because that's just how I think
		for p := p1; true; p = p.Add(ds) {
			grid[p] += 1
			if p == p2 {
				break
			}
		}
	}

	// Find all points in grid greater than 2
	sum := 0

	for _, v := range grid {
		if v > 1 {
			sum++
		}
	}

	return sum
}

// Find the direction of a point that has 0 as one of its values
func direction(p utils.Point) utils.Point {
	magnitude := p.X + p.Y

	if magnitude < 0 {
		magnitude = -magnitude
	}

	return utils.Point{X: p.X / magnitude, Y: p.Y / magnitude}
}

func parseLine(l string) (utils.Point, utils.Point) {
	var p1 utils.Point
	var p2 utils.Point

	_, err := fmt.Sscanf(l, "%d,%d -> %d,%d", &p1.X, &p1.Y, &p2.X, &p2.Y)
	utils.Check(err, "Error parsing line")
	return p1, p2
}
