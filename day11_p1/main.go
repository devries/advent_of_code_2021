package main

import (
	"fmt"
	"io"
	"os"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day11.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	grid := parseGrid(lines)

	flashcount := 0
	for i := 0; i < 100; i++ {
		flashcount += step(grid)
	}

	return flashcount
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

var neighbors = []utils.Point{{X: -1, Y: -1}, {X: -1, Y: 0}, {X: -1, Y: 1}, {X: 0, Y: -1}, {X: 0, Y: 1}, {X: 1, Y: -1}, {X: 1, Y: 0}, {X: 1, Y: 1}}

// Step by one and return number of flashes
func step(grid map[utils.Point]int) int {
	// Increment each by one, saving those that are flashing
	flashers := make([]utils.Point, 0)

	for k, v := range grid {
		grid[k] = v + 1
		if v >= 9 {
			flashers = append(flashers, k)
		}
	}

	// Flash the flashers and and increment neighbors, if there are new flashers then append them
	flashcount := 0
	flashedPoints := make(map[utils.Point]bool) // Track which points have flashed already
	for len(flashers) > 0 {
		p := flashers[0]
		flashers = flashers[1:]
		if flashedPoints[p] {
			continue // We already did this one
		}
		for _, n := range neighbors {
			pn := p.Add(n)
			if flashedPoints[pn] {
				// Already flashed, skip it
				continue
			}
			v, ok := grid[pn]
			if !ok {
				continue // outside grid
			}
			grid[pn] = v + 1
			if v >= 9 {
				flashers = append(flashers, pn)
			}
		}
		grid[p] = 0
		flashedPoints[p] = true
		flashcount++
	}

	// displayGrid(grid)
	return flashcount
}

func displayGrid(grid map[utils.Point]int) {
	i := 0
	j := 0

	for {
		v, ok := grid[utils.Point{X: i, Y: j}]
		if !ok {
			if i == 0 {
				fmt.Printf("\n")
				break
			}
			j++
			i = 0
			fmt.Printf("\n")
			continue
		}
		fmt.Printf("%d", v)
		i++
	}
}
