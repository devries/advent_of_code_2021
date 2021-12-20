package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/devries/advent_of_code_2021/utils"
	"github.com/spf13/pflag"
)

func main() {
	pflag.Parse()
	f, err := os.Open("../inputs/day20.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	r := solve(f, utils.Verbose)
	fmt.Println(r)
}

func solve(r io.Reader, verbose bool) int {
	lines := utils.ReadLines(r)

	key, grid := parseInput(lines)
	for i := 0; i < 2; i++ {
		grid = step(key, grid)
		if verbose {
			fmt.Println(grid)
		}
	}

	return grid.Count()
}

type Grid struct {
	Xmin int
	Xmax int
	Ymin int
	Ymax int
	Data map[utils.Point]int
	Fill int
}

func (g Grid) getPixelSquare(p utils.Point) int {
	ret := 0

	for k, v := range neighbors {
		pt := p.Add(k)
		var pval int

		if pt.X < g.Xmin || pt.X > g.Xmax || pt.Y < g.Ymin || pt.Y > g.Ymax {
			pval = g.Fill
		} else {
			pval = g.Data[pt]
		}

		ret |= pval << v
	}

	return ret
}

func (g Grid) String() string {
	var sb strings.Builder

	for j := g.Ymin; j <= g.Ymax; j++ {
		for i := g.Xmin; i <= g.Xmax; i++ {
			v := g.Data[utils.Point{X: i, Y: j}]
			switch v {
			case 0:
				sb.WriteRune('.')
			case 1:
				sb.WriteRune('#')
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (g Grid) Count() int {
	total := 0

	for _, v := range g.Data {
		total += v
	}

	return total
}

var neighbors = map[utils.Point]int{
	{X: -1, Y: -1}: 8,
	{X: 0, Y: -1}:  7,
	{X: 1, Y: -1}:  6,
	{X: -1, Y: 0}:  5,
	{X: 0, Y: 0}:   4,
	{X: 1, Y: 0}:   3,
	{X: -1, Y: 1}:  2,
	{X: 0, Y: 1}:   1,
	{X: 1, Y: 1}:   0,
}

func parseInput(lines []string) ([]int, Grid) {
	key := make([]int, 0)
	for _, c := range lines[0] {
		switch c {
		case '.':
			key = append(key, 0)
		case '#':
			key = append(key, 1)
		}
	}

	ysize := len(lines) - 2
	xsize := len(lines[2])

	grid := Grid{0, xsize - 1, 0, ysize - 1, make(map[utils.Point]int), 0}
	for j := 0; j < ysize; j++ {
		for i, c := range lines[j+2] {
			if c == '#' {
				grid.Data[utils.Point{X: i, Y: j}] = 1
			}
		}
	}

	return key, grid
}

func step(key []int, grid Grid) Grid {
	newGrid := Grid{grid.Xmin - 1, grid.Xmax + 1, grid.Ymin - 1, grid.Ymax + 1, make(map[utils.Point]int), 0}

	// Handle fill
	if grid.Fill == 0 && key[0] == 1 {
		newGrid.Fill = 1
	} else if grid.Fill == 1 && key[511] == 0 {
		newGrid.Fill = 0
	}

	for j := grid.Ymin - 1; j <= grid.Ymax+1; j++ {
		for i := grid.Xmin - 1; i <= grid.Xmax+1; i++ {
			p := utils.Point{X: i, Y: j}
			square := grid.getPixelSquare(p)
			newGrid.Data[p] = key[square]
		}
	}

	return newGrid
}
