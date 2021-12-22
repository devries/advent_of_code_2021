package main

import (
	"fmt"
	"math"
	"os"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day17.txt")
	utils.Check(err, "error opening input file")
	defer f.Close()

	lines := utils.ReadLines(f)

	var xmin, xmax, ymin, ymax int
	_, err = fmt.Sscanf(lines[0], "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)
	utils.Check(err, "unable to parse input")

	r := solve(xmin, xmax, ymin, ymax)
	fmt.Println(r)
}

func solve(xmin, xmax, ymin, ymax int) int {
	found := make(map[utils.Point]bool)

	for vy := maxVy(ymin); vy > -maxVy(ymin); vy-- {
		stepsA := stepsFromY(ymax, vy)
		stepsB := stepsFromY(ymin, vy)
		for s := stepsA - 1; s <= stepsB+1; s++ {
			if s < 1 {
				continue
			}
			y := ypos(vy, s)
			vxmin := minVx(xmin)
			for vx := vxmin; true; vx++ {
				x := xpos(vx, s)
				if x <= xmax && x >= xmin && y <= ymax && y >= ymin {
					// fmt.Printf("pos: (%d,%d), vx=%d, vy=%d\n", x, ypos(vy, s), vxmin, vy)
					found[utils.Point{X: vx, Y: vy}] = true
				}
				if x > xmax {
					break
				}
			}
		}
	}
	return len(found)
}

func maxVy(y int) int {
	if y > 0 {
		return y
	}
	return -y + 1
}

func minVx(x int) int {
	sq := math.Sqrt(1.0 + 8.0*float64(x))
	return int((sq + 1.0) / 2.0)
}

func ypos(vy, steps int) int {
	return steps*vy - steps*(steps-1)/2
}

func xpos(vx, steps int) int {
	if steps > vx {
		return vx * (vx + 1) / 2
	}
	return (2*vx - steps + 1) * steps / 2
}

func stepsFromY(y, vy int) int {
	b := float64(2*vy + 1)
	srt := math.Sqrt(b*b - 8.0*float64(y))

	sola := (b + srt) / 2.0
	solb := (b - srt) / 2.0

	if solb > 0.0 {
		fmt.Println("Need to consider another domain")
	}

	return int(sola)
}
