package main

import (
	"fmt"
	"io"
	"os"

	"github.com/devries/advent_of_code_2021/utils"
	"github.com/spf13/pflag"
)

func main() {
	pflag.Parse()
	f, err := os.Open("../inputs/day25.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

var down = utils.Point{X: 0, Y: 1}
var right = utils.Point{X: 1, Y: 0}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	width, height, grid := parseInput(lines)
	if utils.Verbose {
		fmt.Print("\x1b[2J")
	}

	steps := 0
	var newGrid map[utils.Point]rune

	for {
		motion := false
		if utils.Verbose {
			fmt.Printf("\033[3J\033[H")
			printGrid(width, height, grid)
		}

		// Iterate
		newGrid = make(map[utils.Point]rune)

		for p, c := range grid {
			if c == '>' {
				destination := p.Add(right)
				if destination.X >= width {
					destination.X = 0
				}
				if grid[destination] == 0 {
					newGrid[destination] = c
					motion = true
				} else {
					newGrid[p] = c
				}
			} else {
				newGrid[p] = c
			}
		}
		grid = newGrid

		newGrid = make(map[utils.Point]rune)

		for p, c := range grid {
			if c == 'v' {
				destination := p.Add(down)
				if destination.Y >= height {
					destination.Y = 0
				}
				if grid[destination] == 0 {
					newGrid[destination] = c
					motion = true
				} else {
					newGrid[p] = c
				}
			} else {
				newGrid[p] = c
			}
		}
		grid = newGrid
		steps++

		if motion == false {
			break
		}
	}
	return steps
}

func parseInput(lines []string) (int, int, map[utils.Point]rune) {
	res := make(map[utils.Point]rune)
	y := len(lines)
	x := len(lines[0])

	for j, ln := range lines {
		for i, c := range ln {
			if c != '.' {
				res[utils.Point{X: i, Y: j}] = c
			}
		}
	}

	return x, y, res
}

func printGrid(xsize int, ysize int, grid map[utils.Point]rune) {
	for j := 0; j < ysize; j++ {
		for i := 0; i < xsize; i++ {
			c := grid[utils.Point{X: i, Y: j}]
			if c == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%c", c)
			}
		}
		fmt.Printf("\n")
	}
}
