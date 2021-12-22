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

func solve(r io.Reader) int64 {
	lines := utils.ReadLines(r)

	inst := parseInput(lines)
	total := int64(0)
	cubes := make([]Cube, len(inst))
	for i := 0; i < len(inst); i++ {
		cubes[i] = inst[i].Region
	}

	for i := 0; i < len(inst); i++ {
		a := inst[i]
		if a.TurnOn {
			total += openPoints(a.Region, cubes[i+1:])
		}
	}

	return total
}

func openPoints(c Cube, blockers []Cube) int64 {
	total := c.Size()

	for i := 0; i < len(blockers); i++ {
		if coverage, ok := c.Intersection(blockers[i]); ok {
			total -= openPoints(coverage, blockers[i+1:])
		}
	}

	return total
}

type Point struct {
	X int
	Y int
	Z int
}

func (p Point) Add(p2 Point) Point {
	return Point{p.X + p2.X, p.Y + p2.Y, p.Z + p2.Z}
}

func (p Point) Sub(p2 Point) Point {
	return Point{p.X - p2.X, p.Y - p2.Y, p.Z - p2.Z}
}

func (p Point) Positive() bool {
	return p.X >= 0 && p.Y >= 0 && p.Z >= 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Cube struct {
	Low  Point
	High Point
}

// Calculate intersection and return boolean if any cubes intersect along with intersection cube
func (c Cube) Intersection(o Cube) (Cube, bool) {
	sectionA := c.High.Sub(o.Low)
	sectionB := o.High.Sub(c.Low)

	if sectionA.Positive() && sectionB.Positive() {
		// intersection
		lowCorner := Point{max(c.Low.X, o.Low.X), max(c.Low.Y, o.Low.Y), max(c.Low.Z, o.Low.Z)}
		highCorner := Point{min(c.High.X, o.High.X), min(c.High.Y, o.High.Y), min(c.High.Z, o.High.Z)}

		return Cube{lowCorner, highCorner}, true
	}

	return Cube{}, false
}

func (c Cube) Size() int64 {
	x := int64(c.High.X-c.Low.X) + 1
	y := int64(c.High.Y-c.Low.Y) + 1
	z := int64(c.High.Z-c.Low.Z) + 1

	return x * y * z
}

type Instruction struct {
	Region Cube
	TurnOn bool
}

func parseInput(lines []string) []Instruction {
	instructions := make([]Instruction, 0)

	for _, line := range lines {
		var low Point
		var high Point
		var command string
		_, err := fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &command, &low.X, &high.X, &low.Y, &high.Y, &low.Z, &high.Z)
		utils.Check(err, "error parsing line")

		instructions = append(instructions, Instruction{Cube{low, high}, command == "on"})
	}

	return instructions
}
