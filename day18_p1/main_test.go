package main

import (
	"strings"
	"testing"
)

var testInput = `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

func TestParser(t *testing.T) {
	tests := []string{
		"[1,2]",
		"[[1,2],3]",
		"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
	}

	for _, test := range tests {
		_, pos := parseSnailNumber([]rune(test), 0)

		if pos != len(test) {
			t.Errorf("Expected %d, got %d", len(test), pos)
		}
	}
}

func TestExplosion(t *testing.T) {
	tests := []struct {
		input  string
		answer string
	}{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	}

	for _, test := range tests {
		el, _ := parseSnailNumber([]rune(test.input), 0)
		explode(el)
		result := el.String()

		if result != test.answer {
			t.Errorf("Expected %s, got %s", test.answer, result)
		}
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		input  string
		answer string
	}{
		{"[[[[0,7],4],[15,[0,13]]],[1,1]]", "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"},
		{"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]"},
	}

	for _, test := range tests {
		el, _ := parseSnailNumber([]rune(test.input), 0)
		split(el)
		result := el.String()

		if result != test.answer {
			t.Errorf("Expected %s, got %s", test.answer, result)
		}
	}
}

func TestSolution(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 4140},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := solve(r, testing.Verbose())

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
