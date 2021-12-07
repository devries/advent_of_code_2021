package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day07.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)
	parts := strings.Split(lines[0], ",")

	values := make([]int, 0)

	for _, p := range parts {
		v, err := strconv.Atoi(p)
		utils.Check(err, "unable to convert string to integer")

		values = append(values, v)
	}

	sort.Ints(values)
	// Search the parameter space

	interval := values[len(values)-1] - values[0]
	cpos := values[0] + interval/2
	interval = interval / 4
	fuels := make(map[int]int)

	for {
		h := cpos + interval
		l := cpos - interval

		cfuel := fuelCalc(values, cpos, fuels)
		hfuel := fuelCalc(values, h, fuels)
		lfuel := fuelCalc(values, l, fuels)

		if interval == 1 && cfuel <= hfuel && cfuel <= lfuel {
			// Success
			return cfuel
		}

		if lfuel <= cfuel && lfuel <= hfuel {
			cpos = l
		} else if hfuel <= cfuel && hfuel <= lfuel {
			cpos = h
		}

		if interval > 1 {
			interval = interval / 2
		}
	}
}

func fuelCalc(values []int, pos int, fuels map[int]int) int {
	if f, ok := fuels[pos]; ok {
		return f
	}

	f := 0
	for _, v := range values {
		d := v - pos
		if d < 0 {
			d = -d
		}
		for i := 1; i <= d; i++ {
			f += i
		}
	}

	fuels[pos] = f
	return f
}
