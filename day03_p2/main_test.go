package main

import (
	"strings"
	"testing"
)

var testInput = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func TestSolution(t *testing.T) {
	tests := []struct {
		input  string
		answer uint64
	}{
		{testInput, 230},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := solve(r)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
