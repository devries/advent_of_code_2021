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
	f, err := os.Open("../inputs/day06.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f, 256)
	fmt.Println(r)
}

func solve(r io.Reader, days int) int64 {
	lines := utils.ReadLines(r)
	parts := strings.Split(lines[0], ",")

	population := make([]int64, 9)

	for _, p := range parts {
		v, err := strconv.Atoi(p)
		utils.Check(err, "unable to convert string to integer")

		population[v] += 1
	}

	for i := 0; i < days; i++ {
		pop0 := population[0]

		// Shift counters down
		for j := 0; j < 8; j++ {
			population[j] = population[j+1]
		}

		// Add babies
		population[8] = pop0

		// Add in reset clocks:
		population[6] += pop0
	}

	sum := int64(0)
	for _, v := range population {
		sum += v
	}
	return sum
}
