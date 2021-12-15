package main

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day15.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	grid := parseGrid(lines)
	total := solveGrid(grid)

	return total
}

type Grid struct {
	Points map[utils.Point]GridPoint
	X      int
	Y      int
}

// Need to define this for easier sorting later
type GridPoint struct {
	Point utils.Point
	Value int
}

func parseGrid(lines []string) Grid {
	xsize := len([]rune(lines[0]))
	ysize := len(lines)
	r := Grid{make(map[utils.Point]GridPoint), xsize * 5, ysize * 5}

	for j := 0; j < r.Y; j++ {
		for i := 0; i < r.X; i++ {
			ln := lines[j%ysize]
			c := []rune(ln)[i%xsize]
			scoreIncrement := i/xsize + j/ysize
			r.Points[utils.Point{X: i, Y: j}] = GridPoint{utils.Point{X: i, Y: j}, (int(c-'0')-1+scoreIncrement)%9 + 1}
		}
	}

	return r
}

type State struct {
	Score int
	Pos   utils.Point
}

func solveGrid(grid Grid) int {
	best := make(map[utils.Point]int) // Best score found to that point
	queue := make([]State, 0)

	start := utils.Point{X: 0, Y: 0}
	best[start] = 0
	queue = append(queue, State{0, start})

	for len(queue) > 0 {
		currentState := queue[0]
		queue = queue[1:]

		// Has this point seen a better outcome already?
		if currentState.Score > best[currentState.Pos] {
			continue
		}

		next := []GridPoint{}
		for _, d := range utils.Directions {
			n, ok := grid.Points[currentState.Pos.Add(d)]
			if ok {
				if b, ok := best[n.Point]; ok {
					vnext := currentState.Score + n.Value
					if b > vnext {
						best[n.Point] = vnext
						next = append(next, n)
					}
				} else {
					best[n.Point] = currentState.Score + n.Value
					next = append(next, n)
				}
			}
		}
		sort.Slice(next, func(i, j int) bool { return next[i].Value < next[j].Value })

		for _, gp := range next {
			queue = append(queue, State{currentState.Score + gp.Value, gp.Point})
		}
	}

	return best[utils.Point{X: grid.X - 1, Y: grid.Y - 1}]
}
