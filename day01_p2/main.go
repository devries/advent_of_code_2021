package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day01.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	// First convert line measurements into sliding 3 point sums
	values := []int{}

	for _, ln := range lines {
		curr, err := strconv.Atoi(ln)
		utils.Check(err, fmt.Sprintf("Error converting %s to int", ln))
		values = append(values, curr)
	}

	sums := []int{}

	for i := 0; i < len(values)-2; i++ {
		sum := values[i] + values[i+1] + values[i+2]
		sums = append(sums, sum)
	}

	increaseCount := 0
	var prev int

	for i, curr := range sums {
		// Convert line to integer, and check if it increased or decreased from previous measurement
		if i > 0 && curr-prev > 0 {
			increaseCount++
		}
		prev = curr
	}

	return increaseCount
}
