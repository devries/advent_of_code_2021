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

	total := 0
	for _, ln := range lines {
		nums, display := parseLine(ln)

		conversion := make(map[uint8]int) // Conversion from binary form of letters to number
		reverse := make(map[int]uint8)    // Also keep a reverse conversion

		for _, s := range nums {
			u := convertString(s)

			switch len(s) {
			case 2:
				conversion[u] = 1
				reverse[1] = u
			case 3:
				conversion[u] = 7
				reverse[7] = u
			case 4:
				conversion[u] = 4
				reverse[4] = u
			case 7:
				conversion[u] = 8
				reverse[8] = u
			}
		}

		for _, s := range nums {
			// Select length 6 characters
			if len(s) == 6 {
				u := convertString(s)

				if u&reverse[4] == reverse[4] {
					conversion[u] = 9
					reverse[9] = u
				} else if u&reverse[1] == reverse[1] {
					conversion[u] = 0
					reverse[0] = u
				} else {
					conversion[u] = 6
					reverse[6] = u
				}
			}
		}

		for _, s := range nums {
			if len(s) == 5 {
				u := convertString(s)

				if u&reverse[1] == reverse[1] {
					conversion[u] = 3
					reverse[3] = u
				} else if u&reverse[6] == u {
					conversion[u] = 5
					reverse[5] = u
				} else {
					conversion[u] = 2
					reverse[2] = u
				}
			}
		}

		displayvalues := make([]int, 0)
		for _, s := range display {
			u := convertString(s)
			displayvalues = append(displayvalues, conversion[u])
		}
		value := displayvalues[0]*1000 + displayvalues[1]*100 + displayvalues[2]*10 + displayvalues[3]
		total += value
	}

	return total
}

func parseLine(s string) ([]string, []string) {
	parts := strings.Split(s, " | ")
	numbers := strings.Fields(parts[0])
	display := strings.Fields(parts[1])

	return numbers, display
}

func convertString(s string) uint8 {
	var output uint8

	for _, r := range s {
		output |= 1 << (r - 'a')
	}

	return output
}
