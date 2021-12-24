package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day24.txt")
	utils.Check(err, "error opening input file")
	defer f.Close()

	solve(f)
}

func solve(r io.Reader) {
	lines := utils.ReadLines(r)

	start := ALU{}
	seen := make(map[ALU]bool)
	ExecuteRecursively(start, lines, []int64{}, seen)
}

type ALU struct {
	// registers map[string]Register
	w  int64
	x  int64
	y  int64
	z  int64
	pc int // program counter
}

func (a *ALU) Set(reg string, v int64) {
	switch reg {
	case "w":
		a.w = v
	case "x":
		a.x = v
	case "y":
		a.y = v
	case "z":
		a.z = v
	}
}

func (a *ALU) Get(reg string) int64 {
	switch reg {
	case "w":
		return a.w
	case "x":
		return a.x
	case "y":
		return a.y
	case "z":
		return a.z
	default:
		v, err := strconv.ParseInt(reg, 10, 64)
		utils.Check(err, fmt.Sprintf("unable to parse %s as number", reg))
		return v
	}
}

func ExecuteRecursively(a ALU, lines []string, digits []int64, seen map[ALU]bool) {
	if len(digits) == 3 {
		fmt.Println(digits)
	}

	for a.pc < len(lines) {
		ln := lines[a.pc]
		parts := strings.Fields(ln)

		switch parts[0] {
		case "inp":
			if seen[a] {
				return
			}
			seen[a] = true

			// a.pc++
			for i := int64(9); i > 0; i-- {
				nd := make([]int64, len(digits)+1)
				copy(nd, digits)
				nd[len(digits)] = i
				n := a
				n.pc++
				n.Set(parts[1], i)
				ExecuteRecursively(n, lines, nd, seen)
			}

			return

		case "add":
			v := a.Get(parts[1]) + a.Get(parts[2])
			a.Set(parts[1], v)

		case "mul":
			v := a.Get(parts[1]) * a.Get(parts[2])
			a.Set(parts[1], v)

		case "div":
			d := a.Get(parts[2])
			if d == 0 {
				return
			}
			v := a.Get(parts[1]) / d
			a.Set(parts[1], v)

		case "mod":
			n := a.Get(parts[1])
			d := a.Get(parts[2])
			if n < 0 || d <= 0 {
				return
			}
			v := n % d
			a.Set(parts[1], v)

		case "eql":
			if a.Get(parts[1]) == a.Get(parts[2]) {
				a.Set(parts[1], 1)
			} else {
				a.Set(parts[1], 0)
			}
		}
		a.pc++
	}

	if a.z == 0 {
		fmt.Println("DONE")
		fmt.Println(digits)
		os.Exit(0)
	}
}
