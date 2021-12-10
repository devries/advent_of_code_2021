package main

import (
	"fmt"
	"io"
	"os"

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

	total := 0
	for _, l := range lines {
		s := NewStack()

	outer:
		for _, r := range l {
			switch r {
			case '(', '[', '{', '<':
				s.Push(r)
			case ')', ']', '}', '>':
				sr, ok := s.Pop()
				if ok == false {
					fmt.Println("Exhausted Stack")
					break outer
				}
				if sr != matching[r] {
					total += points[r]
					break outer
				}
			}
		}
	}

	return total
}

var points = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
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
