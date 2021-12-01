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

	increaseCount := 0

	var prev int

	for i, ln := range lines {
		// Convert line to integer, and check if it increased or decreased from previous measurement
		curr, err := strconv.Atoi(ln)
		utils.Check(err, fmt.Sprintf("Error converting %s to int", ln))
		if i > 0 && curr-prev > 0 {
			increaseCount++
		}
		prev = curr
	}

	return increaseCount
}
