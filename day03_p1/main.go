package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day03.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) uint64 {
	lines := utils.ReadLines(r)

	values := []uint64{}

	for _, ln := range lines {
		v := parseLine(ln)
		values = append(values, v)
	}

	bcheck := uint64(1)
	gamma := uint64(0)
	epsilon := uint64(0)

	for i := 0; i < 12; i++ {
		ones := 0
		zeros := 0
		for _, v := range values {
			br := v & bcheck
			if br == 0 {
				zeros++
			} else {
				ones++
			}
		}
		switch {
		case ones == 0:
		case ones > zeros:
			gamma |= bcheck
		case ones < zeros:
			epsilon |= bcheck
		}
		bcheck <<= 1
	}

	return gamma * epsilon
}

func parseLine(s string) uint64 {
	v, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		panic(err)
	}

	return v
}
