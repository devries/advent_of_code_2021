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
	f, err := os.Open("../inputs/day13.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	solve(f)
	if !utils.Verbose {
		fmt.Println("use -v to see result")
	}
}

func solve(r io.Reader) {
	lines := utils.ReadLines(r)

	pts, folds := parseSheet(lines)
	for _, fold := range folds {
		fold.do(pts)
	}

	if utils.Verbose {
		maxX := 0
		maxY := 0

		for k := range pts {
			if k.X > maxX {
				maxX = k.X
			}
			if k.Y > maxY {
				maxY = k.Y
			}
		}

		for j := 0; j <= maxY; j++ {
			for i := 0; i <= maxX; i++ {
				if pts[utils.Point{X: i, Y: j}] {
					fmt.Printf("\u2588")
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Printf("\n")
		}
	}
}

type Fold struct {
	Orientation rune
	Position    int
}

func (f Fold) do(pts map[utils.Point]bool) {
	foldable := make([]utils.Point, 0)

	switch f.Orientation {
	case 'x':
		for pt := range pts {
			if pt.X > f.Position {
				foldable = append(foldable, pt)
			}
		}

		for _, pt := range foldable {
			delete(pts, pt)
			pt.X = 2*f.Position - pt.X
			pts[pt] = true
		}
	case 'y':
		for pt := range pts {
			if pt.Y > f.Position {
				foldable = append(foldable, pt)
			}
		}

		for _, pt := range foldable {
			delete(pts, pt)
			pt.Y = 2*f.Position - pt.Y
			pts[pt] = true
		}
	}
}

func parseSheet(lines []string) (map[utils.Point]bool, []Fold) {
	pts := make(map[utils.Point]bool)
	folds := make([]Fold, 0)

	for _, ln := range lines {
		switch {
		case len(ln) == 0:
			// Skip blank lines
		case strings.HasPrefix(ln, "fold"):
			// Fold instruction
			var fold Fold
			_, err := fmt.Sscanf(ln, "fold along %c=%d", &fold.Orientation, &fold.Position)
			utils.Check(err, "Unable to parse fold statement")
			folds = append(folds, fold)
		default:
			// Point
			var pt utils.Point
			_, err := fmt.Sscanf(ln, "%d,%d", &pt.X, &pt.Y)
			utils.Check(err, "unable to parse point")
			pts[pt] = true
		}
	}

	return pts, folds
}
