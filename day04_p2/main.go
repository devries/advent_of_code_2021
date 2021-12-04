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
	f, err := os.Open("../inputs/day04.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

// Bingo card
type Card struct {
	Dim     int   // length of row or column
	Numbers []int // Numbers on card
}

func NewCard(d int) *Card {
	c := Card{d, make([]int, d*d)}

	return &c
}

func (c *Card) Parse(lines []string) {
	for i := 0; i < c.Dim; i++ {
		sequence := strings.Fields(lines[i])
		for j, s := range sequence {
			n, err := strconv.Atoi(s)
			utils.Check(err, "Error parsing card number")
			c.Numbers[i*c.Dim+j] = n
		}
	}
}

// Check for Bingo
func (c *Card) Check(selections map[int]bool) bool {
	// Check rows
outerRow:
	for j := 0; j < c.Dim; j++ {
		for i := c.Dim * j; i < c.Dim*(j+1); i++ {
			if selections[c.Numbers[i]] == false {
				continue outerRow
			}
		}
		return true
	}

	// Check columns
outerCol:
	for j := 0; j < c.Dim; j++ {
		for i := j; i < c.Dim*c.Dim; i += c.Dim {
			if selections[c.Numbers[i]] == false {
				continue outerCol
			}
		}
		return true
	}

	return false
}

func (c *Card) Score(selection map[int]bool) int {
	sum := 0
	for _, v := range c.Numbers {
		if selection[v] == false {
			sum += v
		}
	}

	return sum
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	pickLine := strings.Split(lines[0], ",")
	picks := make([]int, 0)

	for _, s := range pickLine {
		v, err := strconv.Atoi(s)
		utils.Check(err, "Error converting integer in pick line")
		picks = append(picks, v)
	}

	cards := make([]*Card, 0)
	for i := 2; i < len(lines); i += 6 {
		c := NewCard(5)
		c.Parse(lines[i : i+5])
		cards = append(cards, c)
	}

	selected := make(map[int]bool)
	eliminated := make([]bool, len(cards)) // True if corresponding card is eliminated
	var lastScore int

	for _, p := range picks {
		selected[p] = true

		for i, c := range cards {
			if !eliminated[i] && c.Check(selected) {
				lastScore = c.Score(selected) * p
				eliminated[i] = true
			}
		}
	}

	return lastScore
}
