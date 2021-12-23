package main

import (
	"strings"
	"testing"
)

var testInput = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`

func TestSolution(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 44169},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := solve(r)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
