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

	width := len(lines[0])
	values := []uint64{}

	for _, ln := range lines {
		v := parseLine(ln)
		values = append(values, v)
	}

	bcheck := uint64(1) << (width - 1)
	oxval := uint64(0)
	coval := uint64(0)
	mask := uint64(0)
	doOx := true // Keep doing ox
	doCo := true // Keep doing co

	for i := 0; i < width; i++ {
		oxOnes := 0
		oxZeros := 0
		coOnes := 0
		coZeros := 0

		nOx := 0 // Number of ox values
		nCo := 0 // Numver of co2 values

		for _, v := range values {
			if doOx && oxval&mask == v&mask {
				nOx++
				br := v & bcheck
				if br == 0 {
					oxZeros++
				} else {
					oxOnes++
				}
			}
			if doCo && coval&mask == v&mask {
				nCo++
				br := v & bcheck
				if br == 0 {
					coZeros++
				} else {
					coOnes++
				}
			}
		}

		// If there was only one match
		if nOx == 1 {
			for _, v := range values {
				if oxval&mask == v&mask {
					oxval = v
				}
			}
			doOx = false
		}

		if nCo == 1 {
			for _, v := range values {
				if coval&mask == v&mask {
					coval = v
				}
			}
			doCo = false
		}

		mask |= bcheck

		// If there are multiple matches add the appropriate bit and continue
		if doOx && oxOnes >= oxZeros {
			oxval |= bcheck
		}

		if doCo && coOnes < coZeros {
			coval |= bcheck
		}

		bcheck >>= 1
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
