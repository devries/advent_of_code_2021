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

	// The median will be the least distance to any point
	sort.Ints(values)

	med := values[len(values)/2]

	fuel := 0

	for _, v := range values {
		d := v - med
		if d < 0 {
			d = -d
		}
		fuel += d
	}

	return fuel
}
