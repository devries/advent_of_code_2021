package main

import (
	"fmt"
	"io"
	"os"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day14.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	polymer, recipies := parseInput(lines)

	for i := 0; i < 10; i++ {
		polymer = step(polymer, recipies)
	}
	ct := polymer.Count()

	var min, max int
	for _, v := range ct {
		if v > max {
			max = v
		}
		if min == 0 || v < min {
			min = v
		}
	}

	return max - min
}

type Polymer struct {
	Pairs map[string]int
	Last  rune
}

// Count number of elements in polymer
func (p Polymer) Count() map[rune]int {
	ct := make(map[rune]int)

	ct[p.Last]++
	for k, v := range p.Pairs {
		elements := []rune(k)
		ct[elements[0]] += v
	}

	return ct
}

type Recipies map[string][]string

func step(p Polymer, r Recipies) Polymer {
	res := Polymer{make(map[string]int), p.Last}

	for k, v := range p.Pairs {
		prods := r[k]
		res.Pairs[prods[0]] += v
		res.Pairs[prods[1]] += v
	}

	return res
}

func parseInput(lines []string) (Polymer, Recipies) {
	p := Polymer{make(map[string]int), 0}
	r := make(Recipies)

	// Get polymer
	pol := []rune(lines[0])
	p.Last = pol[len(pol)-1]
	for i := 0; i < len(pol)-1; i++ {
		pair := string([]rune{pol[i], pol[i+1]})
		p.Pairs[pair] += 1
	}

	// Get recipies
	for i := 2; i < len(lines); i++ {
		var el1, el2, el3 rune
		_, err := fmt.Sscanf(lines[i], "%c%c -> %c", &el1, &el2, &el3)
		utils.Check(err, "error parsing recipe line")

		r[string([]rune{el1, el2})] = []string{string([]rune{el1, el3}), string([]rune{el3, el2})}
	}

	return p, r
}
