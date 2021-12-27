package main

import (
	"fmt"
	"io"
	"os"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day24.txt")
	utils.Check(err, "error opening input file")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int64 {
	lines := utils.ReadLines(r)

	ic := pullConstantsFromInput(lines)

	s := make([]int, 0)
	pairings := make([]Pairing, 0)

	for i := 0; i < 14; i++ {
		if ic[i].toggle == 1 {
			s = append(s, i)
		} else {
			o := s[len(s)-1]
			s = s[:len(s)-1]
			p := Pairing{o, i, ic[o].yadd + ic[i].xadd}
			pairings = append(pairings, p)
		}
	}

	res := make([]int, 14)

	for _, p := range pairings {
		if p.offset < 0 {
			res[p.indexA] = 1 - p.offset
			res[p.indexB] = 1
		} else {
			res[p.indexA] = 1
			res[p.indexB] = 1 + p.offset
		}
	}

	var resint int64
	m := int64(1)
	for i := 13; i >= 0; i-- {
		resint += m * int64(res[i])
		m *= 10
	}

	return resint
}

type IterationConstants struct {
	toggle int
	xadd   int
	yadd   int
}

type Pairing struct {
	indexA int
	indexB int
	offset int
}

func pullConstantsFromInput(lines []string) []IterationConstants {
	ret := make([]IterationConstants, 0)

	for i := 0; i < 14; i++ {
		ic := IterationConstants{}
		fmt.Sscanf(lines[i*18+4], "div z %d", &ic.toggle)
		fmt.Sscanf(lines[i*18+5], "add x %d", &ic.xadd)
		fmt.Sscanf(lines[i*18+15], "add y %d", &ic.yadd)
		ret = append(ret, ic)
	}

	return ret
}
