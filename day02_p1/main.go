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
	f, err := os.Open("../inputs/day02.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	posx := 0
	posy := 0

	for _, ln := range lines {
		dir, dist := parseLine(ln)

		switch dir {
		case "forward":
			posx += dist
		case "down":
			posy += dist
		case "up":
			posy -= dist
		}

	}
	return posx * posy
}

func parseLine(s string) (string, int) {
	parts := strings.Fields(s)

	i, err := strconv.Atoi(parts[1])
	utils.Check(err, "unable to convert integer")

	return parts[0], i
}
