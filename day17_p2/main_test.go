package main

import (
	"testing"
)

var testInput = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestSolution(t *testing.T) {
	tests := []struct {
		xmin   int
		xmax   int
		ymin   int
		ymax   int
		answer int
	}{
		{20, 30, -10, -5, 112},
	}

	for _, test := range tests {
		result := solve(test.xmin, test.xmax, test.ymin, test.ymax)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
