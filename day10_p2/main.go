package main

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/devries/advent_of_code_2021/utils"
)

func main() {
	f, err := os.Open("../inputs/day10.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	r := solve(f)
	fmt.Println(r)
}

func solve(r io.Reader) int {
	lines := utils.ReadLines(r)

	scores := make([]int, 0)
outer:
	for _, l := range lines {
		s := NewStack()

		for _, r := range l {
			switch r {
			case '(', '[', '{', '<':
				s.Push(r)
			case ')', ']', '}', '>':
				sr, ok := s.Pop()
				if ok == false {
					fmt.Println("Exhausted Stack")
					continue outer
				}
				if sr != matching[r] {
					continue outer
				}
			}
		}
		score := 0
		for {
			r, ok := s.Pop()
			if ok == false {
				break
			}

			score = 5*score + points[r]
		}
		scores = append(scores, score)
	}

	// Sort and find middle score
	sort.Ints(scores)

	return scores[len(scores)/2]
}

var points = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

var matching = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

type Stack []rune

func NewStack() *Stack {
	r := make(Stack, 0)

	return &r
}

func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) Pop() (rune, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	r := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]

	return r, true
}
