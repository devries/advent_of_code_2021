package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/devries/advent_of_code_2021/utils"
	"github.com/spf13/pflag"
)

func main() {
	pflag.Parse()
	f, err := os.Open("../inputs/day03.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) uint64 {
	// Read in the binary numbers
	lines := utils.ReadLines(r)

	width := len(lines[0])
	values := []uint64{}

	for _, ln := range lines {
		v := parseLine(ln)
		values = append(values, v)
	}

	oxval := uint64(0)
	coval := uint64(0)

	oxvalues := make([]uint64, len(values))
	copy(oxvalues, values)

	covalues := make([]uint64, len(values))
	copy(covalues, values)

	// Start filtering the numbers from MSb to LSb
	for mask := uint64(1) << (width - 1); mask > 0; mask >>= 1 {
		// Filter the oxvalues
		if len(oxvalues) == 1 {
			oxval = oxvalues[0]
		} else {
			zeros, ones := countBits(oxvalues, mask)
			if ones >= zeros {
				oxval |= mask
			}
			oxvalues = filterValues(oxvalues, oxval, mask)
		}

		// Filter the covalues
		if len(covalues) == 1 {
			coval = covalues[0]
		} else {
			zeros, ones := countBits(covalues, mask)
			if ones < zeros {
				coval |= mask
			}
			covalues = filterValues(covalues, coval, mask)
		}
	}

	if utils.Verbose {
		fmt.Printf("oxygen generator rating: %d\n", oxval)
		fmt.Printf("CO2 scrubber rating: %d\n", coval)
	}

	return oxval * coval
}

func parseLine(s string) uint64 {
	v, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		panic(err)
	}

	return v
}

// Count the number of zeros and ones in position (numbered from 0 = least significant) of an
// array of unsigned integers. Returns number of zeros and number of ones.
func countBits(values []uint64, mask uint64) (int, int) {
	zeros := 0
	ones := 0

	for _, v := range values {
		if v&mask > 0 {
			ones++
		} else {
			zeros++
		}
	}

	return zeros, ones
}

// Filter a set of values whose masked values match the masked match value.
func filterValues(values []uint64, match uint64, mask uint64) []uint64 {
	m := match & mask
	ret := make([]uint64, 0)

	for _, v := range values {
		if m == v&mask {
			ret = append(ret, v)
		}
	}

	return ret
}
