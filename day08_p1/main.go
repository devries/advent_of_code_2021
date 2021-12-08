package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day08.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	count := 0
	for _, ln := range lines {
		_, display := parseLine(ln)

		for _, s := range display {
			if len(s) == 2 || len(s) == 3 || len(s) == 4 || len(s) == 7 {
				count++
			}
		}
	}

	return count
}

func parseLine(s string) ([]string, []string) {
	parts := strings.Split(s, " | ")
	numbers := strings.Fields(parts[0])
	display := strings.Fields(parts[1])

	return numbers, display
}
