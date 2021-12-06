package main

import (
	"strings"
	"testing"
)

var testInput = `3,4,3,1,2`

func TestSolution(t *testing.T) {
	tests := []struct {
		input  string
		days   int
		answer int64
	}{
		{testInput, 18, 26},
		{testInput, 80, 5934},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := solve(r, test.days)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
